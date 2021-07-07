fn main() {
    // create
    // empty string
    let mut s = String::new();

    // with initial content
    // to_string()
    let data = "initial contents";
    let s = data.to_string();
    // String::from()
    let s = String::from("initial contents");

    // update
    // push() & push_str()
    let mut s = String::from("foo");
    s.push(' ');
    s.push_str("bar");

    // + and format!
    {
        let s1 = String::from("hello, ");
        let s2 = String::from("world!");
        // s1不再可用
        // s2使用引用，这里&s2的类型是&String但方法签名中是&str
        // rust 在这里使用了解引用强制多态(deref coercion)即使，类似将&s2转变为了&st[..]
        // 调用约定函数 fn add(self, s: &str) -> String {...}
        let s3 = s1 + &s2;
        println!("{}", s3);
    }
    {
        // 复杂拼接
        let s1 = String::from("tic");
        let s2 = String::from("tac");
        let s3 = String::from("toe");

        let s = s1 + "-" + &s2 + "-" + &s3;
        println!("{}", s);
    }
    {
        // format!宏
        let s1 = String::from("tic");
        let s2 = String::from("tac");
        let s3 = String::from("toe");

        let s = format!("{}-{}-{}", s1, s2, s3);
        println!("{}", s);
    }

    // 索引字符串
    {
        // valid
        let len = String::from("Hola").len();
        println!("{}", len);

        let len = String::from("Здравствуйте").len();
        println!("{}", len);
    }
    {
        // invalid
        let hello = "नमस्ते";
        // 224 返回一个字节
        // let answer = &hello[0];
        // Unicode标量值
        // ['न', 'म', 'स', '्', 'त', 'े']
        // 字形簇
        // ["न", "म", "स्", "ते"]
    }

    // 字符串slice
    {
        let hello = "Здравствуйте";
        let s = &hello[0..4];
        // byte index 1 is not a char boundary; it is inside 'З'
        // let s = &hello[0..1];
        println!("{}", s);
    }

    // 字符串便利
    {
        let hello = "Здравствуйте";
        for c in hello.chars() {
            println!("{}", c);
        }
        for c in hello.bytes() {
            println!("{}", c);
        }
    }
}

