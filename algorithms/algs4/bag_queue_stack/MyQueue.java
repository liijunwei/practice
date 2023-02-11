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
        assertEquals(1, q.size());
    }

    @Test
    public void testEnqueueTwoItems() {
        MyQueue<Integer> q = new MyQueue();
        q.enqueue(2023);
        q.enqueue(2);

        assertFalse(q.isEmpty());
        assertEquals(2, q.size());
        assertEquals(Integer.valueOf(2023), q.first.item);
        assertEquals(Integer.valueOf(2), q.first.next.item);
    }

    @Test
    public void testEnqueueThreeItems() {
        MyQueue<Integer> q = new MyQueue();
        q.enqueue(2023);
        q.enqueue(2);
        q.enqueue(11);

        assertFalse(q.isEmpty());
        assertEquals(3, q.size());
        assertEquals(Integer.valueOf(2023), q.first.item);
        assertEquals(Integer.valueOf(2), q.first.next.item);
        assertEquals(Integer.valueOf(11), q.last.item);
    }

    private int size;
    private Node first;
    private Node last;

    private class Node {
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
        } else {
            Node oldlast = last;
            last = new Node();
            last.item = item;
            last.next = null;
            oldlast.next = last;
        }

        size++;
    }
}
