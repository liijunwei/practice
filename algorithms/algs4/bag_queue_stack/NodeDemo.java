package bag_queue_stack;

import org.junit.Test;

import static org.junit.Assert.*;

public class NodeDemo<Item> {

    @Test
    public void testNodeDemo() {
        Node<String> first = new Node<>();
        Node<String> second = new Node<>();
        Node<String> third = new Node<>();

        first.item = "to";
        second.item = "be";
        third.item = "or";

        first.next = second;
        second.next = third;
        third.next = null;

        assertEquals("to", first.item);
        assertEquals(second, first.next);

        assertEquals("be", second.item);
        assertEquals(third, second.next);

        assertEquals("or", third.item);
        assertNull(third.next);
    }

    private static class Node<Item> {
        Item item;
        Node<Item> next;
    }
}
