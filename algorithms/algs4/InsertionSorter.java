import edu.princeton.cs.algs4.Selection;
import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;
import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class InsertionSorter {
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
            StdOut.println(comparable);
        }
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
