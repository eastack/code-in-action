use std::thread;
use std::time::Duration;

fn main() {
    let v = vec![1, 2, 3];

    let handle = thread::spawn(move || {
        for i in 1..10 {
            println!("Here's a vector: {:?}", v);
            thread::sleep(Duration::from_millis(i))
        }
    });

    handle.join().unwrap();
}
