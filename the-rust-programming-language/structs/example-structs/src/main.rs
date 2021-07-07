fn main() {
    // 多个变量
    {
        let width = 1920;
        let height = 1200;

        fn area(width: u32, height: u32) -> u32 {
            width * height
        }
        println!("The area of the rectangle is {} square pixels.",
                 area(width, height));
    }

    // 一个元组
    {
        let rect = (1920, 1200);

        fn area(dimensions: (u32, u32)) -> u32 {
            dimensions.0 * dimensions.1
        }

        println!("The area of the rectangle is {} square pixels.",
                 area(rect));
    }

    // 结构体
    {
        struct Rectangle {
            width: u32,
            height: u32,
        }

        fn area(rectangle: &Rectangle) -> u32 {
            rectangle.width * rectangle.height
        }

        let rect = Rectangle { width: 1920, height: 1200 };
        println!("The area of the rectangle is {} square pixels.",
                 area(&rect));
    }

    // 派生 trait
    {
        #[derive(Debug)]
        struct Rectangle {
            width: u32,
            height: u32,
        }

        let rect = Rectangle { width: 1920, height: 1200 };
        println!("rect is {:#?}", rect);
    }
}