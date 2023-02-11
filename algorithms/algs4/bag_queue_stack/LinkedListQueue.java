package bag_queue_stack;

import java.util.Iterator;

public class LinkedListQueue<Item> implements Iterable<Item> {
    private Node first; // initial value is null
    private Node last;
    private int size;

    private class Node {
        Item item;
        Node next;
    }

    /**
     * append new item to the last one of the nodes
     *
     */
    public void enqueue(Item item) {
        Node oldlast = last;

        last = new Node();
        last.item = item;
        last.next = null;

        if (isEmpty()) {
            first = last;
        } else {
            oldlast.next = last;
        }

        size++;
    }

    /**
     * return the first item added into the queue
     *
     * @return item
     */
    public Item dequeue() {
        Item item = first.item;
        first = first.next;

        if (isEmpty()) {
            last = null;
        }

        size--;

        return item;
    }

    public boolean isEmpty() {
        return first == null;
    }

    public int size() {
        return size;
    }

    @Override
    public Iterator<Item> iterator() {
        return new ListIterator();
    }

    private class ListIterator implements Iterator<Item> {
        private Node current = first;

        public boolean hasNext() {
            return current != null;
        }

        public void remove() {
        }

        public Item next() {
            Item item = current.item;
            current = current.next;

            return item;
        }
    }

}
