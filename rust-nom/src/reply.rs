use nom::IResult;
use nom::bytes::complete::{tag, take_while, take_while1, take_while_m_n};
use nom::sequence::delimited;
use nom::combinator::map;
use nom::multi::many_m_n;
use nom::branch::alt;
use bytes::BytesMut;
use std::fmt::{Display, Formatter};

#[derive(Debug)]
pub enum Reply {
    // 状态回复或单行回复
    SingleLine(String),
    // 错误回复
    Err(String),
    // 整数回复
    Int(u64),
    // 批量回复
    Batch(Option<String>),
    // 多条批量回复
    MultiBatch(Option<Vec<Reply>>),
    // 回复中没有，这里是为了方便进行错误处理添加的
    BadReply(String),
}

impl Display for Reply {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        match self {
            Reply::SingleLine(line) => write!(f, "+ {}", line),
            Reply::Err(err) => write!(f, ": {}", err),
            Reply::Int(int) => write!(f, ": {}", int),
            Reply::Batch(reply) => {
                if let Some(reply) = reply {
                    write!(f, "$ {}", reply)
                } else {
                    write!(f, "$-1")
                }
            }
            Reply::MultiBatch(replies) => {
                if let Some(replies) = replies {
                    write!(
                        f,
                        "* {}\r\n{}",
                        replies.len(),
                        replies.iter()
                            .map(|r| format!("{}", r))
                            .collect::<Vec<String>>()
                            .join("\r\n")
                    )
                } else {
                    write!(f, "*-1")
                }
            }
            Reply::BadReply(err) => write!(f, "parse reply failed: {}", err)
        }
    }
}

impl Reply {
    pub fn from_resp(src: &BytesMut) -> Self {
        match parse(&String::from_utf8(src.as_ref().to_vec()).unwrap()) {
            Ok((remain, resp)) => {
                if remain.is_empty() {
                    resp
                } else {
                    Reply::BadReply(format!("remaining bytes: {}", remain))
                }
            }
            Err(e) => {
                Reply::BadReply(e.to_string())
            }
        }
    }
}

fn parse_single_line(i: &str) -> IResult<&str, Reply> {
    let (i, _) = tag("+")(i)?;
    let (i, resp) = take_while(|c| c != '\r' && c != '\n')(i)?;
    let (i, _) = tag("\r\n")(i)?;
    Ok((i, Reply::SingleLine(resp.to_string())))
}

fn parse_err(i: &str) -> IResult<&str, Reply> {
    let (i, resp) = delimited(
        tag("-"),
        take_while1(|c| c != '\r' && c != '\n'),
        tag("\r\n"),
    )(i)?;
    Ok((i, Reply::Err(String::from(resp))))
}

fn parse_int(i: &str) -> IResult<&str, Reply> {
    let (i, int) = delimited(
        tag(":"),
        map(take_while1(|c: char| c.is_digit(10)), |int: &str| {
            int.parse::<u64>().unwrap()
        }),
        tag("\r\n"),
    )(i)?;
    Ok((i, Reply::Int(int)))
}

fn parse_batch(i: &str) -> IResult<&str, Reply> {
    let (i, _) = tag("$")(i)?;
    let (i, len) = take_while1(|c: char| c.is_digit(10))(i)?;
    if len == "-1" {
        let (i, _) = tag("\r\n")(i)?;
        Ok((i, Reply::Batch(None)))
    } else {
        let len = len.parse::<usize>().unwrap();
        let (i, resp) = delimited(
            tag("\r\n"),
            take_while_m_n(len, len, |_| true),
            tag("\r\n"),
        )(i)?;
        Ok((i, Reply::Batch(Some(String::from(resp)))))
    }
}

fn parse_multi_batch(i: &str) -> IResult<&str, Reply> {
    let (i, count) = delimited(
        tag("*"),
        take_while1(|c: char| c.is_digit(10)),
        tag("\r\n"),
    )(i)?;

    if count == "-1" {
        let (i, _) = tag("\r\n")(i)?;
        Ok((i, Reply::MultiBatch(None)))
    } else {
        let count = count.parse::<usize>().unwrap();
        let (i, responses) = many_m_n(
            count,
            count,
            alt((parse_single_line, parse_err, parse_int, parse_batch)),
        )(i)?;
        if responses.len() != count {
            Ok((i, Reply::BadReply(format!("expect {} items, get {}", count, responses.len()))))
        } else {
            Ok((i, Reply::MultiBatch(Some(responses))))
        }
    }
}

pub fn parse(i: &str) -> IResult<&str, Reply> {
    alt((
        parse_single_line,
        parse_err,
        parse_int,
        parse_batch,
        parse_multi_batch
    ))(i)
}

