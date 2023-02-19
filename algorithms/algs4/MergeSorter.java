import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class MergeSorter {
    private static Comparable[] aux;

    @Test
    public void testInsertionSorter() {
        Comparable[] list1 = {1, 3, 2, 5, 8, 4, 6, 0};
        Comparable[] list2 = {0, 1, 2, 3, 4, 5, 6, 8};

        MergeSorter.sort(list1);

        assertEquals(list2, list1);
    }

    private static void sort(Comparable[] list1) {
    }

    public static void merge(Comparable[] a, int lo, int mid, int hi) {
        int i = lo; // 左半边的cursor
        int j = mid + 1; // 右半边的cursor

        for (int k = lo; k <= hi; k++) {
            aux[k] = a[k]; // 将a[lo..hi] 复制到aux[lo..hi]
        }

        for (int k = lo; k <= hi; k++) {
            if (i > mid) { // 左半边用尽，取右半边的元素
                a[k] = aux[j];
                j++;
            } else if (j > hi) { // 右半边用尽，取左半边的元素
                a[k] = aux[i];
                i++;
            } else if (less(aux[j], aux[i])) { // 右半边当前元素小于左半边的当前元素，取右半边的元素
                a[k] = aux[j];
                j++;
            } else { // 右半边当前元素大于等于左半边当前元素，取左半边的元素
                a[k] = aux[i];
                i++;
            }
        }
    }

    private static boolean less(Comparable v, Comparable w) {
        return v.compareTo(w) < 0;
    }
}
