public class MaxPQArray<Key extends Comparable<Key>> {
    private Key[] pq;
    private int size = 0;

    public MaxPQArray(int maxN) {
        pq = (Key[]) new Comparable[maxN + 1];
    }

    public boolean isEmpty() {
        return size == 0;
    }

    public int size() {
        return size;
    }

    public void insert(Key v) {
        pq[size++] = v;
    }

    public Key delMax() {
        int maxIndex = 0;

        for (int i = maxIndex + 1; i < size; i++) {
            if (less(pq[maxIndex], pq[i])) {
                maxIndex = i;
            }
        }

        Key max = pq[maxIndex];
        exch(pq, maxIndex, size - 1);
        size--;

        return max;
    }

    public Key max() {
        return null;
    }

    private boolean less(Comparable v, Comparable w) {
        return v.compareTo(w) < 0;
    }

    private void exch(Object[] a, int i, int j) {
        Object tmp = a[i];
        a[i] = a[j];
        a[j] = tmp;
    }
}
