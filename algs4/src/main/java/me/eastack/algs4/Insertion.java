package me.eastack.algs4;

import edu.princeton.cs.algs4.StdOut;

@SuppressWarnings("rawtypes, unchecked")
public class Insertion {
    public static void oldSort(Comparable[] a) {
        int N = a.length;
        for (int i = 0; i < N; i++) {
            for (int j = i; j > 0 && less(a[j], a[j - 1]); j--) {
                exch(a, j, j - 1);
            }
        }
    }

    public static void sort(Comparable[] a) {
        int N = a.length;
        for (int i = 0; i < N; i++) {

            int j = i;
            Comparable temp = a[j];

            while (j > 0 && less(temp, a[j - 1])) {
                a[j] = a[j - 1];
                j--;
            }

            a[j] = temp;
        }
    }

    private static boolean less(Comparable v, Comparable w) {
        return v.compareTo(w) < 0;
    }

    private static void exch(Comparable[] a, int i, int j) {
        Comparable t = a[i];
        a[i] = a[j];
        a[j] = t;
    }

    private static void show(Comparable[] a) {
        for (Comparable comparable : a) {
            StdOut.print(comparable + " ");
        }
        StdOut.println();
    }

    public static boolean isSorted(Comparable[] a) {
        for (int i = 1; i < a.length; i++) {
            if (less(a[i], a[i - 1])) return false;
        }
        return true;
    }

    public static void main(String[] args) {
        Comparable[] arr = new Comparable[]{
            1, 6, 3, 5, 6, 8, 9, 0, 5, 67, 3, 2, 1, 6, 7, 8, 134, 567, 2, 2, 542456, 256, 25, 2435, 2345, 234
        };

        sort(arr);
        if (isSorted(arr)) show(arr);

        Comparable[] arr2 = new Comparable[]{
            1, 6, 3, 5, 6, 8, 9, 0, 5, 67, 3, 2, 1, 6, 7, 8, 134, 567, 2, 2, 542456, 256, 25, 2435, 2345, 234
        };

        oldSort(arr2);
        if (isSorted(arr2)) show(arr);
    }
}
