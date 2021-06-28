use nom::error::{ErrorKind, ParseError, ContextError, context};
use std::collections::HashMap;
use nom::{IResult, InputTakeAtPosition, AsChar};
use nom::bytes::complete::{take_while, tag, escaped};
use nom::sequence::{delimited, preceded, terminated, separated_pairc, separated_pair};
use nom::branch::alt;
use nom::combinator::{map, opt, cut, value};
use nom::character::complete::{char, one_of, alphanumeric1, alphanumeric0};
use nom::multi::separated_list0;
use nom::number::complete::double;
use std::fs;
use nom::character::streaming::alpha1;

fn main() {
    let data = "  { \"a\"\t: 42,
  \"b\": [ \"x\", \"y\", 12 ] ,
  \"c\": { \"hello\" : \"world\"
  }
  } ";

    let json = fs::read_to_string("data.json").unwrap();

    println!(
        "will try to parse valid JSON data:\n\n**********\n{}\n**********\n",
        data
    );

    println!("parsing a valid file:\n{:#?}\n",
             root::<(&json, ErrorKind)>(&json));
}

#[derive(Debug, PartialEq)]
pub enum JsonValue {
    Str(String),
    Boolean(bool),
    Null,
    Num(f64),
    Array(Vec<JsonValue>),
    Object(HashMap<String, JsonValue>),
}

fn root<'a, E: ParseError<&'a str> + ContextError<&'a str>>(
    i: &'a str,
) -> IResult<&'a str, JsonValue, E> {
    delimited(
        sp,
        alt((map(hash, JsonValue::Object),
             map(array, JsonValue::Array))),
        opt(sp),
    )(i)
}

fn sp<'a, E: ParseError<&'a str> + ContextError<&'a str>>(
    i: &'a str,
) -> IResult<&'a str, &'a str, E> {
    let chars = " \t\r\n";
    take_while(move |c| chars.contains(c))(i)
}


fn hash<'a, E: ParseError<&'a str> + ContextError<&'a str>>(
    i: &'a str,
) -> IResult<&'a str, HashMap<String, JsonValue>, E> {
    context(
        "map",
        // 扔第一个，拿第二个
        preceded(
            // 消费一个 {
            char('{'),
            // 根据条件剪切
            cut(
                // 拿第一个，扔掉第二个
                terminated(
                    // 匹配并转换
                    map(
                        // 交替匹配分隔符和元素
                        separated_list0(preceded(sp, char(',')), key_value),
                        |tuple_vec| {
                            tuple_vec
                                .into_iter()
                                .map(|(k, v)| (String::from(k), v))
                                .collect()
                        },
                    ),
                    preceded(sp, char('}')),
                )),
        ),
    )(i)
}

fn key_value<'a, E: ParseError<&'a str> + ContextError<&'a str>>(
    i: &'a str,
) -> IResult<&'a str, (&'a str, JsonValue), E> {
    separated_pair(
        preceded(sp, string),
        cut(preceded(sp, char(':'))),
        json_value,
    )(i)
}

fn json_value<'a, E: ParseError<&'a str> + ContextError<&'a str>>(
    i: &'a str,
) -> IResult<&'a str, JsonValue, E> {
    preceded(
        sp,
        alt((
            map(hash, JsonValue::Object),
            map(array, JsonValue::Array),
            map(string, |s| JsonValue::Str(String::from(s))),
            map(double, JsonValue::Num),
            map(boolean, JsonValue::Boolean),
        )),
    )(i)
}

fn array<'a, E: ParseError<&'a str> + ContextError<&'a str>>(
    i: &'a str,
) -> IResult<&'a str, Vec<JsonValue>, E> {
    context(
        "array",
        preceded(
            char('['),
            cut(terminated(
                separated_list0(preceded(sp, char(',')), json_value),
                preceded(sp, char(']')),
            )),
        ),
    )(i)
}

fn boolean<'a, E: ParseError<&'a str> + ContextError<&'a str>>(
    i: &'a str,
) -> IResult<&'a str, bool, E> {
    let parse_true = value(true, tag("true"));
    let parse_false = value(false, tag("false"));
    alt((parse_true, parse_false))(i)
}

fn string<'a, E: ParseError<&'a str> + ContextError<&'a str>>(
    i: &'a str,
) -> IResult<&'a str, &'a str, E> {
    context(
        "string",
        preceded(
            char('\"'),
            cut(terminated(parse_str, char('\"')))),
    )(i)
}

fn parse_str<'a, E: ParseError<&'a str>>(i: &'a str) -> IResult<&'a str, &'a str, E> {
    escaped(alt((alphanumeric1, underline)),
            '\\',
            one_of("\"n\\"))(i)
}

pub fn underline<T, E: ParseError<T>>(input: T) -> IResult<T, T, E>
    where
        T: InputTakeAtPosition,
        <T as InputTakeAtPosition>::Item: AsChar,
{
    input.split_at_position1_complete(|item| item.as_char() != '_', ErrorKind::Alpha)
}