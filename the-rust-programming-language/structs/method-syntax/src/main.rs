fn main() {
    #[derive(Debug)]
    struct Rectangle {
        width: u32,
        height: u32,
    }

    // impl 块可以分散
    impl Rectangle {
        // 参数self参与自动引用和解引用
        fn area(&self) -> u32 {
            self.width * self.height
        }
    }
    impl Rectangle {
        // 多参数方法
        fn can_hold(&self, other: &Rectangle) -> bool {
            self.width > other.width && self.height > other.height
        }

        // 关联函数
        fn square(size: u32) -> Rectangle {
            Rectangle { width: size, height: size }
        }
    }

    let rect0 = Rectangle { width: 1920, height: 1200 };
    // 关联函数调用使用 :: 语法
    let rect1 = Rectangle::square(3000);
    println!("The area of the rectangle is {} square pixels.",
             rect0.area());

    println!("Can rect0 hold rect1? {}", rect0.can_hold(&rect1));
}