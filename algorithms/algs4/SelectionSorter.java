import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class SelectionSorter {
    public static void sort(Comparable[] a) {
        int N = a.length;

        for (int i = 0; i < N; i++) {
            int minIndex = i;

            for (int j = minIndex + 1; j < N; j++) {
                if (less(a[j], a[minIndex])) {
                    minIndex = j;
                }
            }

            swap(a, i, minIndex);
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
    public void testSelectionSorter() {
        String[] list1 = {"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"};
        String[] list2 = {"A", "E", "E", "L", "M", "O", "P", "R", "S", "T", "X"};

        SelectionSorter.sort(list1);

        assertEquals(list2, list1);
    }
}
