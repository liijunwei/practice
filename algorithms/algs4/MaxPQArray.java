/**
 * 2.4.2.1 MaxPQ array implementation
 * @param <Key>
 */
public class MaxPQArray<Key extends Comparable<Key>> {
    private Key[] pq;
    private int size = 0;

    public MaxPQArray(int maxN) {
        pq = (Key[]) new Comparable[maxN];
    }

    public boolean isEmpty() {
        return size == 0;
    }

    public int size() {
        return size;
    }

    public int arrLength() {
        return pq.length;
    }

    public void insert(Key v) {
        if (size == pq.length) {
            resize(2 * pq.length);
        }

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

        if (size > 0 && size == (pq.length / 4)) {
            resize(pq.length / 2);
        }

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

    private void resize(int max) {
        Key[] temp = (Key[]) new Comparable[max];

        if (size >= 0) {
            System.arraycopy(pq, 0, temp, 0, size);
        }

        pq = temp;
    }
}
