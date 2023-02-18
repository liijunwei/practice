import org.junit.Test;

import static org.junit.Assert.assertEquals;

//https://www.sortvisualizer.com/shellsort/
public class ShellSorter {
    public static void sort(Comparable[] a) {
        int n = a.length;
        int h = 3;

        while (h < n / 3) {
            h = 3 * h + 1; // 1, 4, 13, 40, 121, 364
        }

        /**
         * 如果我们在插入排序（算法2。2）中加入一个外循环来将h 按照递增序列递减
         * 我们就能得到这个简洁的希尔排序
         */
        while (h >= 1) {
            for (int i = h; i < n; i++) {
                Comparable tmp = a[i];
                int j = i;

                for (; j >= h && less(tmp, a[j - h]); j -= h) {
                    a[j] = a[j - h];
                }

                a[j] = tmp;
            }

            h = h / 3;
        }
    }

    private static void swap(Object[] a, int i, int j) {
        Object tmp = a[i];
        a[i] = a[j];
        a[j] = tmp;
    }

    private static boolean less(Comparable v, Comparable w) {
        return v.compareTo(w) < 0;
    }

    @Test
    public void testInsertionSorter() {
        Comparable[] list1 = {1, 3, 2, 5, 8, 4, 6, 0};
        Comparable[] list2 = {0, 1, 2, 3, 4, 5, 6, 8};

        ShellSorter.sort(list1);

        assertEquals(list2, list1);
    }
}
