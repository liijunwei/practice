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

    @Test
    public void testInsertFirstNode() {
        Node<String> first = new Node<>();
        Node<String> second = new Node<>();
        Node<String> third = new Node<>();

        first.item = "to";
        second.item = "be";
        third.item = "or";

        first.next = second;
        second.next = third;
        third.next = null;

        Node<String> oldfirst = first;
        first = new Node<>();
        first.item = "inserted first";
        first.next = oldfirst;

        assertEquals("inserted first", first.item);
        assertEquals("to", first.next.item);
        assertEquals(second, first.next.next);
    }

    private static class Node<Item> {
        Item item;
        Node<Item> next;
    }
}
