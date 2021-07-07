use std::collections::HashMap;

fn main() {
    let numbers = vec![1, 2, 3, 4, 5, 5];
    let number_len = numbers.len() as i32;

    let mut sum: i32 = 0;
    for number in &numbers {
        sum = sum + number;
    }

    println!("sum: {}", sum);
    println!("avg: {}", sum / &number_len);

    let mut number_count_map = HashMap::new();
    for number in &numbers {
        let count = number_count_map.entry(number).or_insert(0);
        *count += 1;
    }

    let mut count_number_map = HashMap::new();
    let mut counter: Vec<&i32> = vec![];
    for (k, v) in &number_count_map {
        &counter.push(v);
        count_number_map.insert(v, k);
    }

    println!("counter: {:?}", counter);
    println!("counter max: {:?}", counter.iter().max());
    println!("counter number map: {:?}", count_number_map);

    match counter.iter().max() {
        Some(max) => println!("Max count number: {:?}", count_number_map.get(max)),
        None => (),
    }

    println!("{:?}", number_count_map);
}
