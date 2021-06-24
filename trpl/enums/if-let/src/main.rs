fn main() {
    let some_u8_value = Some(0u8);

    // match
    match some_u8_value {
        Some(3) => println!("three"),
        _ => (),
    }

    // if let
    if let Some(3) = some_u8_value {
        println!("three");
    } else {
        println!("not three");
    }
}
