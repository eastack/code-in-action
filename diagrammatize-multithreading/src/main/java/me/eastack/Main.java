package me.eastack;

public class Main {
    public static void main(String[] args) {
        test();
    }

    public static synchronized void test() {
        System.out.println("a");
        System.out.println(Thread.holdsLock(Main.class));
    }
}
