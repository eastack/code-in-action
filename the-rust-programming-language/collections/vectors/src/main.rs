fn main() {
    // let v: Vec<i32> = Vec::new();
    let mut v = vec![1, 2, 3];
    v.push(4);
    v.push(5);
    v.push(6);
    v.push(7);

    let third = &v[2];
    // v.push(1);
    println!("The third element is {}", third);
    v.push(8);

    match v.get(2) {
        Some(third) => println!("The third element is {}", third),
        None => println!("There is no third element."),
    }
    for i in &mut v {
        *i += 50;
    }

    for i in &v {
        println!("{}", i);
    }

    #[derive(Debug)]
    enum SpreadsheetCell {
        Int(i32),
        Float(f64),
        Text(String),
    }

    let row = vec![
        SpreadsheetCell::Int(3),
        SpreadsheetCell::Text(String::from("blue")),
        SpreadsheetCell::Float(10.12),
    ];

    for c in &row {
        println!("{:?}", c);
    }
}
