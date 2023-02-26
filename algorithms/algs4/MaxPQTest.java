import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class MaxPQTest {
    @Test
    public void testMaxPQArraySize() {
        MaxPQArray<Integer> pq = new MaxPQArray<>(10);
        pq.insert(1);
        pq.insert(10);
        pq.insert(2);
        assertEquals(3, pq.size());
    }
}
