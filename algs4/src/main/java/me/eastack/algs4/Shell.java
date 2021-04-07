package me.eastack.algs4;

@SuppressWarnings("rawtypes")
public class Shell implements Sorter {
    public void sortWithExch(Comparable[] a) {

        int N = a.length;
        int h = 1;

        // 确定递增序列最大值
        while (h < N / 3) {
            h = 3 * h + 1;
        }

        // 外层循环将数组按照递增序列为间隙分组
        while (h >= 1) {
            // 之前的插入排序
            for (int i = h; i < N; i++) {
                for (int j = i;
                     j >= h && less(a[j], a[j - h]);
                     j -= h) {

                    exch(a, j, j - h);
                }
            }
            h /= 3;
        }
    }

    public void sort(Comparable[] a) {

        int N = a.length;
        int h = 1;

        // 确定递增序列最大值
        while (h < N / 3) {
            h = 3 * h + 1;
        }

        // 外层循环将数组按照递增序列为间隙分组
        while (h >= 1) {
            // 之前的插入排序
            for (int i = h; i < N; i++) {

                int j = i;
                Comparable temp = a[j];

                while (j >= h && less(temp, a[j - h])) {
                    a[j] = a[j - h];
                    j = j - h;
                }

                a[j] = temp;

            }
            h = h / 3;
        }
    }

    public static void main(String[] args) {
        new Shell().sort();
    }
}
