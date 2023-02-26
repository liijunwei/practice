import bag_queue_stack.Stack;

/**
 * 2.4.2.3 MaxPQ linked list implementation
 *
 * @param <Key>
 */
public class MaxPQLinkedList<Key extends Comparable<Key>> {
    private Node first;
    private int size;

    private class Node {
        Key item;
        Node next;
    }

    /**
     * same as stack push
     *
     * @param item new item to insert
     */
    public void insert(Key item) {
        Node oldfirst = first;
        first = new Node();
        first.item = item;
        first.next = oldfirst;
        size++;
    }

    /**
     * similar to stack pop, but it find the max item to return
     * 暂时想不明白
     * @return max item
     */
    public Key delMax() {
        Node maxPrev = first;

        while (less(maxPrev.item, maxPrev.next.item)) {
            maxPrev = maxPrev.next;
        }

        Node max = maxPrev.next;
        Key item = max.item;
        maxPrev.next = max.next;
        size--;

        return item;
    }

    private boolean less(Comparable v, Comparable w) {
        return v.compareTo(w) < 0;
    }
}
