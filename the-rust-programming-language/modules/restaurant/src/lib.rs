mod front_of_house;

// 可以使用 pub use 进行重新导出
// 常用于代码构建者和使用者对问题域关注点不同，导致不同的认知，所以导出的模块也相应有不同结构
pub use crate::front_of_house::hosting;

fn server_order() {}

mod back_of_house {
    pub enum Appetizer {
        Soup,
        Salad,
    }

    pub struct Breakfast {
        pub toast: String,
        seasonal_fruit: String,
    }

    impl Breakfast {
        pub fn summer(toast: &str) -> Breakfast {
            Breakfast {
                toast: String::from(toast),
                seasonal_fruit: String::from("peaches"),
            }
        }
    }

    fn fix_incorrect_order() {
        cook_order();
        super::server_order();
    }

    fn cook_order() {}
}

pub fn eat_at_restaurant() {
    // Absolute path
    // crate::front_of_house::hosting::add_to_waitlist();

    // Relative path
    // front_of_house::hosting::add_to_waitlist();

    let mut meal = back_of_house::Breakfast::summer("Rye");
    meal.toast = String::from("Wheat");
    // meal.seasonal_fruit = String::from("blueberries");
    println!("I'd like {} toast please", meal.toast);

    let order1 = back_of_house::Appetizer::Soup;
    let order2 = back_of_house::Appetizer::Salad;

    // 惯用use路径，只是一种习惯，便于我们写出易于阅读的代码
    // 对于函数，我们为了清楚区分调用函数声明位置，所以携带部分模块名
    use crate::front_of_house::hosting;

    hosting::add_to_waitlist();
    hosting::add_to_waitlist();
    hosting::add_to_waitlist();
    // 对于数据结构，我们直接将其引入作用域
    use std::collections::HashMap;

    // 可以使用 as 进行重命名
    use std::collections::HashMap as Map;


    // 可以使用路径嵌套消除大量 use 行
    // use std::cmp::Ordering;
    // use std::io;
    // use std::{cmp::Ordering, io};
    // use std::io::{self, Write};

    // 使用 glob 运算符，谨慎使用
    // use std::collections::*;
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
