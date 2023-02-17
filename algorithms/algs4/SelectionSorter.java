public class SelectionSorter {
    public static void sort(Comparable[] a) {
        int N = a.length;

        for (int i = 0; i < N; i++) {
            int minIndex = i;

            for (int j = minIndex + 1; j < N; j++) {
                if (less(a[j], a[minIndex])) {
                    minIndex = j;
                }

                swap(a, i, j);
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
}
