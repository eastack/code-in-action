fn main() {
    {
        let number_list = vec![34, 45, 90, 100];

        let mut largest = number_list[0];

        for number in number_list {
            if number > largest {
                largest = number;
            }
        }

        println!("The largest number is {}", largest);


        let number_list = vec![10, 3, 6, 2, 9];
        let mut largest = number_list[0];

        for number in number_list {
            if number > largest {
                largest = number;
            }
        }

        println!("The largest number is {}", largest);
    }
    {
        let number_list = vec![10, 3, 6, 2, 9];
        println!("{}", largest(&number_list));
        let number_list = vec![34, 45, 90, 100];
        println!("{}", largest(&number_list));
        fn largest(list: &[i32]) -> i32 {
            let mut largest = list[0];

            for &item in list {
                if item > largest {
                    largest = item;
                }
            }

            largest
        }

    }
}



