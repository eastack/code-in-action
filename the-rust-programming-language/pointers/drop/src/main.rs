struct CustomSmartPointer {
    data: String,
}

impl Drop for CustomSmartPointer {
    fn drop(&mut self) {
        println!("Dropping CustomSmartPointer with data `{}`!", self.data);
    }
}

fn main() {
    let c = CustomSmartPointer {data: String::from("my stuff")};
    println!("CustomSmartPointer created.");
    // error 无法显示调用drop，不然最后就会double free了
    // c.drop();
    drop(c);
    let d = CustomSmartPointer {data: String::from("other stuff")};
    println!("CustomSmartPointers created.")
}
