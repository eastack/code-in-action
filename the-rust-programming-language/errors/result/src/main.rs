use std::fs::File;
use std::io;
use std::io::{ErrorKind, Read};

fn main() {
    let f = File::open("hello.txt").expect("Failed to open hello.txt");

    // let f = match f {
    //     Ok(file) => file,
    //     Err(error) => match error.kind() {
    //         ErrorKind::NotFound => match File::create("hello.txt"){
    //             Ok(fc)  => fc,
    //             Err(e) => panic!("Problem creating the file: {:?}", e),
    //         },
    //         other_error => panic!("Problem opening the file: {:?}", error),
    //     },
    // };

    println!("{:?}", f);
}

fn read_username_from_file() -> Result<String, io::Error> {
    let f = File::open("hello.txt");
    let mut f = match f {
        Ok(file) => file,
        // 这里显示的提前返回Err
        Err(e) => return Err(e),
    };

    let mut s = String::new();

    match f.read_to_string(&mut s) {
        Ok(_) => Ok(s),
        // 这里是函数最后一个表达式，所以不需要 return
        Err(e) => Err(e),
    }
}

fn read_username_from_file_with_question() -> Result<String, io::Error> {
    let mut f = File::open("hello.txt")?;
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}

fn read_username_from_file_with_question_chain() -> Result<String, io::Error> {
    let mut s = String::new();
    File::open("hello.txt")?.read_to_string(&mut s)?;
    Ok(s)
}


use std::fs;
fn oh() -> Result<String, io::Error> {
    fs::read_to_string("hello.txt")
}
