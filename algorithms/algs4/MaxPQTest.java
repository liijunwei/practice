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
}
