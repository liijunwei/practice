import edu.princeton.cs.algs4.StdRandom;
import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class QuickSorter {
    @Test
    public void testQuickSorter() {
        Comparable[] list1 = {1, 3, 2, 5, 8, 4, 6, 0};
        Comparable[] list2 = {0, 1, 2, 3, 4, 5, 6, 8};

        QuickSorter.sort(list1);

        assertEquals(list2, list1);
    }

    public static void sort(Comparable[] a) {
        StdRandom.shuffle(a);
        sort(a, 0, a.length - 1);
    }

    private static void sort(Comparable[] a, int lo, int hi) {
        if (hi <= lo) {
            return;
        }

        int j = partition(a, lo, hi);
        sort(a, lo, j - 1);
        sort(a, j + 1, hi);
    }

    private static int partition(Comparable[] a, int lo, int hi) {
        int i = lo;
        int j = hi + 1; // ???
        Comparable v = a[lo];

        while (true) {
            // 如果左指针 或者右指针 碰到了边界，就不再移动

            // 将左指针i 向右移动，直至 a[i] >= 切割点的值
            while (less(a[++i], v)) {
                if (i == hi) {
                    break;
                }
            }

            // 将右指针j 向左移动，直至 a[j] <= 切割点的值
            while (less(v, a[--j])) {
                if (j == lo) {
                    break;
                }
            }

            if (i >= j) {
                break;
            }

            // i/j 找好位置以后，交换值，然后进入下一轮循环
            // 直到循环退出
            // 然后把 切割点放到合适的位置
            swap(a, i, j);
        }

        swap(a, lo, j); // put v to right position

        return j;
    }

    private static boolean less(Comparable v, Comparable w) {
        return v.compareTo(w) < 0;
    }

    private static void swap(Object[] a, int i, int j) {
        Object tmp = a[i];
        a[i] = a[j];
        a[j] = tmp;
    }

}
