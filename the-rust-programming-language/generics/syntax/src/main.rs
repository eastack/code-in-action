fn main() {
    // generic in function
    {
        fn largest_i32(list: &[i32]) -> i32 {
            let mut largest = list[0];

            for &item in list.iter() {
                if item > largest {
                    largest = item;
                }
            }

            largest
        }

        fn largest_char(list: &[char]) -> char {
            let mut largest = list[0];

            for &item in list.iter() {
                if item > largest {
                    largest = item;
                }
            }

            largest
        }

        // need generic support
        use std::cmp::PartialOrd;
        use std::marker::Copy;
        fn largest<T: PartialOrd + Copy>(list: &[T]) -> T {
            let mut largest = list[0];

            for &item in list.iter() {
                if item > largest {
                    largest = item;
                }
            }

            largest
        }

        fn largest2<T: PartialOrd>(list: &[T]) -> usize {
            let mut largest_index = 0;

            for (i, &item) in list.iter().enumerate() {
                if item > list[largest_index] {
                    largest_index = i;
                }
            }

            largest_index
        }

        let numbers = vec![1, 2, 3, 4, 5];
        let largest = numbers[largest2(&numbers)];
        println!("{}", largest)

    }

    // in struct
    {
        struct Point<T> {
            x: T,
            y: T,
        }

        let integer = Point { x: 5, y: 10 };
        let float = Point { x: 5.0, y: 10.0 };
    }
    {
        struct Point<T, U> {
            x: T,
            y: U,
        }

        let both_integer = Point { x: 5, y: 10 };
        let both_float = Point { x: 1.0, y: 4.0 };
        let integer_and_float = Point { x: 1, y: 4.0 };
    }

    // in enum
    {
        enum Option<T> {
            Some(T),
            None,
        }

        enum Result<T, E> {
            Ok(T),
            Err(E),
        }
    }

    // in impl
    {
        struct Point<T> {
            x: T,
            y: T,
        }

        impl<T> Point<T> {
            fn x(&self) -> &T {
                &self.x
            }
        }

        impl Point<f32> {
            fn distance_from_origin(&self) -> f32 {
                (self.x.powi(2) + self.y.powi(2)).sqrt()
            }
        }

        let p = Point { x: 5, y: 10 };
        let fp = Point { x: 5.0 as f32, y: 10.2 as f32 };

        println!("Distance from origin: {}", fp.distance_from_origin());

        println!("p.x = {}", p.x());
    }
    // mix up generic in impl
    {
        struct Point<T, U> {
            x: T,
            y: U,
        }

        impl<T, U> Point<T, U> {
            fn mixup<V, W>(self, other: Point<V, W>) -> Point<T, W> {
                Point {
                    x: self.x,
                    y: other.y,
                }
            }
        }

        let p1 = Point { x: 5, y: 10.4 };
        let p2 = Point { x: "Hello", y: "c" };

        let p3 = p1.mixup(p2);

        println!("p3.x = {}, p3.y = {}", p3.x, p3.y);
    }
}

