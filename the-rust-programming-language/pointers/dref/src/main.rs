use std::ops::Deref;

fn main() {
    let x = 5;
    let y = &x;

    let x = 5;
    let y = MyBox::new(x);

    assert_eq!(5, x);
    assert_eq!(5, *(y.deref()));

    let m = MyBox::new(String::from("Rust"));
    // hello(&(*m)[..]);
    // 相比实际的方式，这里rust自动完成了（deref coercions）解引用强制转换
    hello(&m);
}

struct MyBox<T>(T);

impl<T> MyBox<T> {
    fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}

impl<T> Deref for MyBox<T> {
    type Target = T;

    fn deref(&self) -> &T {
        &self.0
    }
}

fn hello(name: &str) {
    println!("Hello, {}!", name);
}