import org.junit.Ignore;
import org.junit.Test;

import static org.junit.Assert.*;

public class MaxPQTest {
    @Test
    public void testMaxPQArrayQueryMethods() {
        MaxPQArray<Integer> pq = new MaxPQArray<>(10);
        assertTrue(pq.isEmpty());

        pq.insert(1);
        pq.insert(10);
        pq.insert(2);
        assertFalse(pq.isEmpty());
        assertEquals(3, pq.size());
    }

    @Test
    public void testMaxPQArrayDeleteMax() {
        MaxPQArray<Integer> pq = new MaxPQArray<>(10);
        pq.insert(1);
        pq.insert(10);
        pq.insert(2);

        assertEquals(Integer.valueOf(10), pq.delMax());
        assertEquals(Integer.valueOf(2), pq.delMax());

        pq.insert(2000);
        assertEquals(Integer.valueOf(2000), pq.delMax());

        pq.insert(100);
        assertEquals(Integer.valueOf(100), pq.delMax());
    }

    @Test
    public void testMaxPQArrayResizeOnInsertion() {
        MaxPQArray<Integer> pq = new MaxPQArray<>(2);
        assertEquals(0, pq.size());

        pq.insert(1);
        pq.insert(10);
        pq.insert(2);

        assertEquals(3, pq.size());
        pq.insert(2);
        pq.insert(2);
        pq.insert(2);
        pq.insert(2);
        assertEquals(7, pq.size());
    }

    @Test
    public void testMaxPQArrayResizeOnDeletion() {
        MaxPQArray<Integer> pq = new MaxPQArray<>(4);
        assertEquals(4, pq.arrLength());

        pq.insert(1);
        pq.insert(10);

        assertEquals(4, pq.arrLength());

        pq.delMax();
        assertEquals(2, pq.arrLength());
    }

    @Test
    public void testMaxPQSortedArray() {
        MaxPQSortedArray<Integer> pq = new MaxPQSortedArray<>(10);
        assertTrue(pq.isEmpty());

        pq.insert(1);
        pq.insert(10);
        pq.insert(2);
        pq.insert(200);
        assertEquals(4, pq.size());

        assertEquals(Integer.valueOf(200), pq.delMax());
        assertEquals(Integer.valueOf(10), pq.delMax());
        assertEquals(Integer.valueOf(2), pq.delMax());
        assertEquals(Integer.valueOf(1), pq.delMax());
    }

    @Ignore
    @Test
    public void testMaxPQLinkedList() {
        MaxPQLinkedList<Integer> pq = new MaxPQLinkedList<>();

        pq.insert(1);
        pq.insert(10);
        pq.insert(2);
        pq.insert(200);

        assertEquals(Integer.valueOf(200), pq.delMax());
        assertEquals(Integer.valueOf(10), pq.delMax());
        assertEquals(Integer.valueOf(2), pq.delMax());
        assertEquals(Integer.valueOf(1), pq.delMax());
    }

    @Test
    public void testMaxPQHeap() {
        MaxPQHeap<Integer> pq = new MaxPQHeap<>(10);

        pq.insert(1);
        pq.insert(10);
        pq.insert(2);
        pq.insert(200);
        pq.insert(15);
        pq.insert(20);
        pq.insert(2000);

        assertEquals(Integer.valueOf(2000), pq.delMax());
        assertEquals(Integer.valueOf(200), pq.delMax());
        assertEquals(Integer.valueOf(20), pq.delMax());
        assertEquals(Integer.valueOf(15), pq.delMax());
        assertEquals(Integer.valueOf(10), pq.delMax());
        assertEquals(Integer.valueOf(2), pq.delMax());
        assertEquals(Integer.valueOf(1), pq.delMax());
    }

    @Test
    public void testMaxPQHeap1() {
        MaxPQHeap<String> pq = new MaxPQHeap<>(10);

        pq.insert("P");
        pq.insert("Q");
        pq.insert("E");

        assertEquals("Q", pq.delMax());

        pq.insert("X");
        pq.insert("A");
        pq.insert("M");

        assertEquals("X", pq.delMax());

        pq.insert("P");
        pq.insert("L");
        pq.insert("E");

        assertEquals("P", pq.delMax());
    }
}
