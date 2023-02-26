/**
 * 2.4.2.2 MaxPQ sorted array implementation
 *
 * @param <Key>
 */
public class MaxPQSortedArray<Key extends Comparable<Key>> {
    private Key[] pq;
    private int size = 0;

    public MaxPQSortedArray(int maxN) {
        pq = (Key[]) new Comparable[maxN];
    }

    public boolean isEmpty() {
        return size == 0;
    }

    public int size() {
        return size;
    }

    public void insert(Key v) {
        if (size == pq.length) {
            resize(2 * pq.length);
        }

        // 插入排序的逻辑: 找到插入位置，之后的元素右移
        // v 是新拿到手里的牌，从后向前扫描，如果发现有更小的元素，就把牌依次向右移动，为新牌腾出位置
        int j = size;

        for (; j > 0 && less(pq[j - 1], v); j--) {
            pq[j] = pq[j - 1];
        }

        pq[j] = v;
        size++;
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
