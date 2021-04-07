package me.eastack.algs4;

@SuppressWarnings("rawtypes")
public class Insertion implements Sorter {
    /**
     * 有交换插入排序
     * @param a 输入数组
     */
    private void ordinarySort(Comparable[] a) {
        int N = a.length;
        for (int i = 0; i < N; i++) {
            for (int j = i; j > 0 && less(a[j], a[j - 1]); j--) {
                exch(a, j, j - 1);
            }
        }
    }

    /**
     * 无交换插入排序
     * @param a 输入数组
     */
    public void sort(Comparable[] a) {
        int N = a.length;
        for (int i = 0; i < N; i++) {

            int j = i;
            Comparable temp = a[j];

            while (j >= 1 && less(temp, a[j - 1])) {
                a[j] = a[j - 1];
                j = j -1;
            }

            a[j] = temp;
        }
    }

    public static void main(String[] args) {
        new Insertion().sort();
    }
}
