use std::collections::HashMap;

fn main() {
    let mut scores = HashMap::new();

    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);

    let teams = vec![String::from("Blue"), String::from("Yellow")];
    let init_scores = vec![10, 50];

    // HashMap类型信息必须有，但反省可以使用下划线省略，让rust自动推断完成
    let scores: HashMap<_, _> = teams.iter()
        .zip(init_scores.iter())
        .collect();

    // 所有权相关
    {
        let field_name = String::from("Favorite color");
        let field_value = String::from("Blue");

        let mut map = HashMap::new();
        // 此处所有权move到map中
        // map.insert(field_name, field_value);
        // field_name, field_value不再有效

        // 如果使用引用则不会move所有权
        map.insert(&field_name, &field_value);

        println!("{}, {}", field_name, field_value);
    }

    // 访问hashmap中的值
    {
        let mut scores = HashMap::new();

        scores.insert(String::from("Blue"), 10);
        scores.insert(String::from("Yellow"), 50);

        let team_name = String::from("Blue");
        let score = scores.get(&team_name);
        match score {
            Some(v) => println!("{}", v),
            None => ()
        }

        if let Some(v) = score {
            println!("{}", v);
        }

        // 顺序不可靠
        for (key, value) in &scores {
            println!("{}: {}", key, value);
        }
    }

    // 更新hashmap
    {
        // 覆盖掉
        let mut scores = HashMap::new();
        scores.insert(String::from("Blue"), 10);
        scores.insert(String::from("Blue"), 50);

        println!("{:?}", scores);

        // insert if not exists
        scores.entry(String::from("Yellow")).or_insert(60);
        scores.entry(String::from("Blue")).or_insert(60);

        println!("{:?}", scores);

        // 根据旧值进行更新
        let text = "hello world wonderful world";
        let mut map = HashMap::new();

        for word in text.split_whitespace() {
            let count = map.entry(word).or_insert(0);
            *count += 1;
        }

        println!("{:?}", map);
    }
}
