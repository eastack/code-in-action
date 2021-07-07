fn main() {
    // 枚举和结构体结合使用
    {
        enum IpAddrKind {
            V4,
            V6,
        }

        let four = IpAddrKind::V4;
        let six = IpAddrKind::V6;

        struct IpAddr {
            kind: IpAddrKind,
            address: String,
        }

        let home = IpAddr {
            kind: IpAddrKind::V4,
            address: String::from("127.0.0.1"),
        };

        let loopback = IpAddr {
            kind: IpAddrKind::V6,
            address: String::from("::1"),
        };
    }

    // 包含数据的枚举
    {
        enum IpAddr {
            V4(String),
            V6(String),
        }

        let home = IpAddr::V4(String::from("127.0.0.1"));
        let loopback = IpAddr::V6(String::from("::1"));
    }

    // 每个枚举成员包含不同数据
    {
        enum IpAddr {
            V4(u8, u8, u8, u8),
            V6(String),
        }

        let home = IpAddr::V4(127, 0, 0, 1);
        let home = IpAddr::V6(String::from("::1"));
    }

    // enum成员丰富类型
    {
        enum Message {
            Quit,
            // 没有任何关联数据
            Move { x: i32, y: i32 },
            // 包含一个匿名结构体
            Write(String),
            // 包含一个单独的String
            ChangeColor(i32, i32, i32), // 包含三个i32
        }

        impl Message {
            fn call(&self) {
                // 在这里定义方法体
                match self {
                    Message::Write(str) => println!("Write: {}" ,str),
                    _ => {}
                }
            }
        }

        let m = Message::Write(String::from("hello"));
        m.call();

        // 上边的枚举和分别定义结构体很像
        // 类单元结构体
        struct QuitMessage;
        struct MoveMessage {
            x: i32,
            y: i32,
        }
        // 元组结构体
        struct WriteMessage(String);
        // 元组结构体
        struct ChangeColorMessage(i32, i32, i32);
    }

    // 让我们看看Rust内置的Optional枚举
    {
        let some_number = Some(5);
        let some_thing =Some("a string");

        // 当为None时需要提供类型信息
        let absent_number: Option<i32> = None;
    }
}
