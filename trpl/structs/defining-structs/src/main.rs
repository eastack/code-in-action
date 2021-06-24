fn main() {
    // 结构体定义
    struct User {
        username: String,
        email: String,
        sign_in_account: u64,
        active: bool,
    }

    {
        // 结构体实例化
        let user = User {
            username: String::from("EaStack"),
            active: true,
            email: String::from("admin@eastack.me"),
            sign_in_account: 1,
        };
    }

    {
        // 修改结构体的值(Rust不允许单字段可变)
        let mut user = User {
            email: String::from("admin@eastack.me"),
            username: String::from("EaStack"),
            active: true,
            sign_in_account: 1,
        };
        user.email = String::from("example@example.com");
    }

    {
        // 结构体初始化语句，是个表达式
        fn build_user(email: String, username: String) -> User {
            User {
                email: email,
                username: username,
                active: true,
                sign_in_account: 1,
            }
        }
    }

    {
        // 结构体初始化
        // 技巧 1 变量与字段同名时的字段初始化简写语法
        fn build_user(email: String, username: String) -> User {
            User {
                email,
                username,
                active: true,
                sign_in_account: 1,
            }
        }
        // 技巧 2 使用结构体更新语法(struct update syntax)从其他实例创建实例
        let user = build_user(String::from("admin@eastck.me"),
                              String::from("EaStack"));

        let user2 = User {
            email: String::from("example@example.com"),
            username: String::from("example"),
            ..user
        };
    }

    {
        // 元组结构体
        struct Color(i32, i32, i32);
        struct Point(i32, i32, i32);

        let black = Color(0, 0, 0);
        let point = Point(0, 0, 0);
    }

    {
        // 类单元结构体/unit-like structs (因为它们类似于 ()，即 unit 类型）
        struct U;
        let u1 = U {};
    }

    println!("Hello, world!");
}
