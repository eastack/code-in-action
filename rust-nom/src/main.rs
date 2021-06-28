mod commands;
mod reply;

use std::error::Error;
use clap::Clap;
use bytes::{BytesMut, BufMut};
use tokio::net::TcpStream;
use tokio::io::{AsyncWriteExt, AsyncReadExt};

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let command = commands::Commands::parse();

    let mut stream = TcpStream::connect("127.0.0.1:6379").await?;
    let mut buf = [0u8; 1024];
    let mut resp = BytesMut::with_capacity(1024);

    let (mut reader, mut writer) = stream.split();

    writer.write(&command.to_bytes()).await?;
    let n = reader.read(&mut buf).await?;
    resp.put(&buf[0..n]);


    let rep = reply::Reply::from_resp(&resp);
    println!("{}", rep);

    Ok(())
}
