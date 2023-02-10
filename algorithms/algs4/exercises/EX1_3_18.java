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
    private static class Node {
        Integer num;
        Node next;
    }

    public static Node findSecondLastNode(Node first) {
        if (first == null || first.next == null) {
            return first;
        }

        Node secondLast = first;

        while (secondLast.next.next != null) {
            secondLast = secondLast.next;
        }
        return secondLast.next;
    }

    @Test
    public void testZeroOrOneNode() {
        Node first = null;
        assertNull(findSecondLastNode(first));

        first = new Node();
        first.num = 1;

        assertEquals(first, findSecondLastNode(first));
    }

    @Test
    public void testTwoOrMoreNodes() {
        Node first = new Node();
        first.num = 1;

        Node second = new Node();
        second.num = 2;

        first.next = second;

        assertEquals(second, findSecondLastNode(first));

        Node third = new Node();
        third.num = 3;

        second.next = third;

        assertEquals(third, findSecondLastNode(first));
    }

}
