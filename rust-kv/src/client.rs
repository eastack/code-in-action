mod pb;

use prost::Message;

use crate::pb::RequestGet;

fn main() {
    let request = RequestGet{
        key: "Hello".to_string()
    };
    let mut buf = Vec::new();
    request.encode(&mut buf).unwrap();
    println!("encoded: {:?}", buf);
}
