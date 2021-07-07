fn main() {
    let mut str = String::from("first");

    let target = 'i';

    if let Some(index) = str.find(target) {
        str.remove(index);
        let result = format!("{}-{}ay", str, target);
        println!("{}", result);
    }
}
