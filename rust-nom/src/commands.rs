use clap::{AppSettings, Clap};
use bytes::{BytesMut, BufMut};

/// redis client cli
#[derive(Clap, Debug)]
#[clap(name = "rcli")]
#[clap(setting = AppSettings::ColoredHelp)]
pub enum Commands {
    /// push value to list
    Rpush {
        /// redis key
        key: String,
        /// value
        values: Vec<String>,
    },
    Set {
        /// redis key
        key: String,

        /// redis key value
        value: String,
    },
    /// get string value
    Get {
        key: String
    },
    /// increase 1
    Incr {
        /// redis key
        key: String,
    },
}

impl Commands {
    pub fn to_bytes(&self) -> bytes::BytesMut {
        match self {
            Commands::Rpush { key, values } => {
                let mut builder = CmdBuilder::new().arg("RPUSH").arg(key);
                values.iter().for_each(|v| builder.add_arg(v));
                builder.to_bytes()
            },
            Commands::Get{key} => {
                let builder =  CmdBuilder::new().arg("GET").arg(key);
                builder.to_bytes()
            },
            _ => {
                println!("todo");
                CmdBuilder::new().to_bytes()
            }
        }
    }
}

struct CmdBuilder {
    args: Vec<String>,
}

impl CmdBuilder {
    fn new() -> Self {
        CmdBuilder { args: vec![] }
    }

    fn arg(mut self, arg: &str) -> Self {
        self.args.push(format!("${}", arg.len()));
        self.args.push(arg.to_string());
        self
    }

    fn add_arg(&mut self, arg: &str) {
        self.args.push(format!("${}", arg.len()));
        self.args.push(arg.to_string());
    }

    fn to_bytes(&self) -> BytesMut {
        let mut bytes = BytesMut::new();
        bytes.put(&format!("*{}\r\n", self.args.len() / 2).into_bytes()[..]);
        bytes.put(&self.args.join("\r\n").into_bytes()[..]);
        bytes.put(&b"\r\n"[..]);
        bytes
    }
}