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
        int i = lo;
        int j = mid + 1;

        for (int k = lo; k <= hi; k++) {
            aux[k] = a[k]; // 将a[lo..hi] 复制到aux[lo..hi]
        }

        for (int k = lo; k <= hi; k++) {
            if (i > mid) {
                a[k] = aux[j];
                j++;
            } else if (j > hi) {
                a[k] = aux[i];
                i++;
            } else if (less(aux[i], aux[j])) {
                a[k] = aux[j];
                j++;
            } else {
                a[k] = aux[i];
                i++;
            }
        }
    }

    private static boolean less(Comparable v, Comparable w) {
        return v.compareTo(w) < 0;
    }
}
