import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class MaxPQArray<Key extends Comparable<Key>> {
    private Key[] pq;
    private int N = 0;

    public MaxPQArray(int maxN) {
        pq = (Key[]) new Comparable[maxN + 1];
    }

    public boolean isEmpty() {
        return N == 0;
    }

    public int size() {
        return N;
    }

    public void insert(Key v) {
        pq[++N] = v;
    }

    public Key delMax() {
        return null;
    }

    public Key max() {
        return null;
    }

    @Test
    public void testMaxPQArrayImplementation() {
        MaxPQArray<Integer> pq = new MaxPQArray<>(10);
        pq.insert(1);
        pq.insert(10);
        pq.insert(2);
//        assertEquals(list2, list1);
    }
}
