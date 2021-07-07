fn main() {
    // 裸指针
    {
        // 允许忽略借用规则，可以同时拥有可变和不可变指针，或多个指向相同位置的可变指针
        // 不保证指向有效的内存
        // 允许为空
        // 不能实现任何自动清理
        let mut num = 5;

        let r1 = &num as *const i32;
        let r2 = &mut num as &mut i32;
    }
}
