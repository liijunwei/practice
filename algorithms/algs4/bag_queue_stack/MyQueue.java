package bag_queue_stack;

import org.junit.Test;

import static org.junit.Assert.*;

public class MyQueue<Item> {
    @Test
    public void testQueueIsEmptyByDefault() {
        MyQueue q = new MyQueue();
        assertTrue(q.isEmpty());
        assertEquals(0, q.size());
    }

    @Test
    public void testEnqueueAnItem() {
        MyQueue<Integer> q = new MyQueue();
        Integer magicNum = 7;
        q.enqueue(magicNum);
        assertFalse(q.isEmpty());
//        assertEquals(1, q.size());
    }

    private int size;
    private Node first;
    private Node last;

    private static class Node<Item> {
        Item item;
        Node next;
    }

    public boolean isEmpty() {
        return size == 0;
    }

    public int size() {
        return size;
    }

    public void enqueue(Item item) {
        if (isEmpty()) {
            first = last = new Node();
            first.item = item;
        }
        size++;
    }
}
