package exercises;

import org.junit.Test;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNull;

/**
 * 1. find the tail node
 * 2. delete the tail node
 * 3. test pass
 */
public class EX1_3_18 {
    private class Node {
        Integer num;
        Node next;
    }

    public static Node removeLast(Node first) {
        if (first == null || first.next == null) {
            return first;
        }

        Node tmp = first;

        while (tmp.next.next != null) {
            tmp = tmp.next;
        }
        return tmp.next;
    }

    @Test
    public void testZeroOrOneNode() {
        Node first = null;
        assertNull(removeLast(first));

        Node n = new Node();
        n.num = 1;
        first = n;

        assertEquals(n, removeLast(first));
    }

    @Test
    public void testTwoOrMoreNodes() {
        Node first = new Node();
        first.num = 1;

        Node second = new Node();
        second.num = 2;

        first.next = second;

        assertEquals(second, removeLast(first));

        Node third = new Node();
        third.num = 3;

        second.next = third;

        assertEquals(third, removeLast(first));
    }

}
