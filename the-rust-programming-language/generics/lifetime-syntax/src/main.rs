use std::fmt::Display;

fn main() {
    {
        // let r;
        // {
        //     let x = 5;
        //     r = &x;
        // }
        //
        // println!("{}", r);
    }
    {
        let x = 5;

        let r = &x;

        println!("{}", r);
    }

    {
        let str1 = String::from("abcd");
        let str2 = "xyz";

        let result = longest(str1.as_str(), str2);

        println!("The longest string is {}", result);
    }

    {
        let str1 = String::from("abcd");
        let result;
        {
            let str2 = String::from("xyz");
            result = longest(str1.as_str(), str2.as_str());
            println!("The longest string is {}", result);
        }
        // println!("The longest string is {}", result);
    }
    {
        struct ImportantExcerpt<'a> {
            part: &'a str,
        }

        let novel = String::from("Call me Ishmael. Some years ago...");
        let first_sentence = novel.split('.').next().expect("Could not find a '.'");
        let i = ImportantExcerpt {
            part: first_sentence,
        };
    }
}

fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn longest_with_an_announcement<'a, T>(x: &'a str, y: &'a str, ann: T) -> &'a str
where
    T: Display,
{
    println!("Announcement! {}", ann);
    if x.len() > y.len() {
        x
    } else {
        y
    }
}
