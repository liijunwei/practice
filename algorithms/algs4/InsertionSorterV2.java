import edu.princeton.cs.algs4.Selection;
import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;
import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class InsertionSorterV2 {
    /**
     * page158
     * "要大幅提高插入排序的速度并不难，只需要在内循环中将较大的元素都向右移动而不总是交换两个元素(这样访问数组的次数就能减半)"
     * see: exercise 2.1.25
     * 右移: a[j] = a[j-1]
     *
     * 先移动数组，再把新的元素放到合适的位置
     */
    public static void sort(Comparable[] a) {
        for (int i = 1; i < a.length; i++) {
            Comparable tmp = a[i];
            int j = i;

            for (; j > 0 && less(tmp, a[j - 1]); j--) {
                a[j] = a[j - 1];
            }

            a[j] = tmp;
            show(a);
        }
    }

    private static boolean less(Comparable v, Comparable w) {
        return v.compareTo(w) < 0;
    }

    private static void show(Comparable[] a) {
        for (Comparable comparable : a) {
            StdOut.print(comparable + " ");
        }
        StdOut.println();
    }

    @Test
    public void testInsertionSorterV2() {
        String[] list1 = {"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"};
        String[] list2 = {"A", "E", "E", "L", "M", "O", "P", "R", "S", "T", "X"};

        InsertionSorterV2.sort(list1);

        assertEquals(list2, list1);
    }

    public static void main(String[] args) {
        String[] a = StdIn.readAllStrings();
        InsertionSorterV2.sort(a);
    }
}
