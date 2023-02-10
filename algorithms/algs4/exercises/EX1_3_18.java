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

    public static String nodeStr(Node first) {
        StringBuilder s = new StringBuilder();

        for (Node temp = first; temp != null; temp = temp.next){
            s.append(temp.num).append(' ');
        }

        return s.toString().trim();
    }

    public static Node removeLastNode(Node first) {
        if (first == null || first.next == null) {
            return null;
        }

        Node secondLast = first;

        while (secondLast.next.next != null) {
            secondLast = secondLast.next;
        }
        secondLast.next = null;

        return first;
    }

    @Test
    public void testZeroOrOneNode() {
        Node first = null;
        assertNull(findSecondLastNode(first));

        first = new Node();
        first.num = 1;

        assertEquals(first, findSecondLastNode(first));
        assertEquals("1", nodeStr(first));

        removeLastNode(first);
        assertEquals("1", nodeStr(first));
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
        assertEquals("1 2 3", nodeStr(first));

        removeLastNode(first);
        assertEquals("1 2", nodeStr(first));
    }

}
