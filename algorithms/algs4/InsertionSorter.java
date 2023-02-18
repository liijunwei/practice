import edu.princeton.cs.algs4.Selection;
import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;
import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class InsertionSorter {
    /**
     * 把这种排序方式想像成给发到手的牌 排序
     * 循环里i增大，相当于得到一张新的牌
     * 循环里j减小并且比较，相当于给手里的牌重新排序(把刚刚得到的牌插入到合适的位置)
     * <p>
     * 所谓的插入排序，不是一下子把某个数插入到合适的位置，而是指最终到效果看起来好像是把数字插入到了合适到问题
     * 实际上还是遍历和交换到结果
     */
    public static void sort(Comparable[] a) {
        for (int i = 1; i < a.length; i++) {
            for (int j = i; j > 0 && less(a[j], a[j - 1]); j--) {
                swap(a, j, j - 1);
            }
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

    private static void show(Comparable[] a) {
        for (Comparable comparable : a) {
            StdOut.print(comparable + " ");
        }
        StdOut.println();
    }

    @Test
    public void testInsertionSorter() {
        String[] list1 = {"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"};
        String[] list2 = {"A", "E", "E", "L", "M", "O", "P", "R", "S", "T", "X"};

        InsertionSorter.sort(list1);

        assertEquals(list2, list1);
    }

    public static void main(String[] args) {
        String[] a = StdIn.readAllStrings();
        Selection.sort(a);
        show(a);
    }
}
