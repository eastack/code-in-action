package me.eastack.algs4.ch02.code;

import edu.princeton.cs.algs4.StdOut;

@SuppressWarnings("rawtypes, unchecked")
public interface Sorter {
    void sort(Comparable[] a);

    default boolean less(Comparable v, Comparable w) {
        return v.compareTo(w) < 0;
    }

    default void exch(Comparable[] a, int i, int j) {
        Comparable t = a[i];
        a[i] = a[j];
        a[j] = t;
    }

    default void show(Comparable[] a) {
        for (Comparable comparable : a) {
            StdOut.print(comparable + " ");
        }
        StdOut.println();
    }

    default boolean isSorted(Comparable[] a) {
        for (int i = 1; i < a.length; i++) {
            if (less(a[i], a[i - 1])) {
                return false;
            }
        }
        return true;
    }

    default void sort() {
        Double[] array = SortCompare.randomArray();
        sort(array);
        System.out.printf("Sorted: %b\n", isSorted(array));
    }
}
