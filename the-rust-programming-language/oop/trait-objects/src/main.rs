fn main() {
    let v1 = vec![1, 2, 3];
    let total: i32 = v1.iter().sum();
    println!("{:?}", total);

    let v2: Vec<_> = v1.iter()
        .map(|x| x + 1)
        .collect();

    let total: i32 = v2.iter().sum();
    println!("{}", total);
}
