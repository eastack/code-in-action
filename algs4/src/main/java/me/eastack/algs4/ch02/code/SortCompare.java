package me.eastack.algs4.ch02.code;

import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.StdRandom;
import edu.princeton.cs.algs4.Stopwatch;

import java.util.Locale;
import java.util.function.Function;

public class SortCompare {
    public static double time(String alg, Double[] a) {
        Stopwatch timer = new Stopwatch();
        if (alg.equals("Insertion")) new Insertion().sort(a);
        if (alg.equals("Selection")) new Selection().sort(a);
        return timer.elapsedTime();
    }

    public static Double[] randomArray() {
        return randomArray(1000000);
    }

    public static Double[] randomArray(int len) {
        Double[] result = new Double[len];

        for (int i = 0; i < len; i++) {
            result[i] = StdRandom.uniform();
        }

        return result;
    }

    public static double timeRandomInput(String alg, int N, int T) {
        double total = 0.0;
        for (int t = 0; t < T; t++) {
            total += time(alg, randomArray(N));
        }
        return total;
    }

    public static void main(String[] args) {
        String alg1 = "Insertion";
        String alg2 = "Selection";

        int N = Integer.parseInt("100");
        int T = Integer.parseInt("10");

        double t1 = timeRandomInput(alg1, N, T); // 算法1的总时间
        double t2 = timeRandomInput(alg2, N, T); // 算法1的总时间

        StdOut.printf("For %d random Doubles\n    %s is", N, alg1);
        StdOut.printf(" %.1f times faster than %s\n", t2/t1, alg2);

        Function<String, Integer> c = String::length;
        Function<String, String> b = str -> str.toLowerCase(Locale.ROOT);
        Function<String, Integer> stringIntegerFunction = b.andThen(c);

        System.out.println(stringIntegerFunction.apply("abc"));
    }
}
