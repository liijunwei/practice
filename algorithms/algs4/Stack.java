import java.util.Iterator;

public class Stack<Item> implements Iterable<Item> {
    private Node first;
    private int size;

    public Item peek() {
        return first.item;
    }

    private class Node {
        Item item;
        Node next;
    }

    public void push(Item item) {
        Node oldfirst = first;
        first = new Node();
        first.item = item;
        first.next = oldfirst;
        size++;
    }

    public Item pop() {
        Item item = first.item;
        first = first.next;
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
