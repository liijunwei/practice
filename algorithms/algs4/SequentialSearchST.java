import edu.princeton.cs.algs4.Queue;

import java.util.ArrayList;
import java.util.List;

public class SequentialSearchST<Key, Value> {
    private Node first;
    private final List<Integer> compareTimes = new ArrayList();

    public List<Integer> getComparisonTimes() {
        return compareTimes;
    }

    private class Node {
        Key key;
        Value value;
        Node next;

        public Node(Key key, Value value, Node next) {
            this.key = key;
            this.value = value;
            this.next = next;
        }
    }

    public Value get(Key key) {
        if (key == null) {
            throw new IllegalArgumentException("argument to contains() is null");
        }

        for (Node x = first; x != null; x = x.next) {
            if (key.equals(x.key)) {
                return x.value;
            }
        }

        return null;
    }

    public boolean contains(Key key) {
        if (key == null) {
            throw new IllegalArgumentException("argument to contains() is null");
        }

        return get(key) != null;
    }

    public void put(Key key, Value value) {
        int compareTime = 0;
        if (key == null) {
            compareTimes.add(1);
            throw new IllegalArgumentException("argument to contains() is null");
        }

        for (Node x = first; x != null; x = x.next) {
            compareTime++;

            if (key.equals(x.key)) {
                x.value = value;
                compareTimes.add(compareTime);
            }
        }

        first = new Node(key, value, first); // 未命中，将新节点插入链表头

        compareTimes.add(compareTime);
    }

    // TODO
    public int size() {
        return 0;
    }

    public Iterable<Key> keys() {
        Queue<Key> queue = new Queue<Key>();
        for (Node x = first; x != null; x = x.next)
            queue.enqueue(x.key);
        return queue;
    }

    // TODO
    public void delete(Key key) {

    }
}
