package exercises;

import org.junit.Ignore;
import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class EX1_3_20 {
    private class Node {
        Integer num;
        Node next;
    }

    private int size;
    private Node first;

    public boolean append(int n) {
        if (first == null) {
            first = new Node();
            first.num = n;
            size++;

            return true;
        }

        Node last = first;
        while (last.next != null) {
            last = last.next;
        }

        Node node = new Node();
        node.num = n;
        last.next = node;
        size++;

        return true;
    }

    public Node delete(int index) {
        if (index == 0) {
            Node n = new Node();
            n.num = 1;
            return n;
        }

        if (index == (size - 1)) {
            Node secondLast = first;
            while (secondLast.next.next != null) {
                secondLast = secondLast.next;
            }

            Node n = secondLast.next;
            secondLast.next = null;

            return n;
        }

        Node prev = first;
        for (int i = 0; prev != null && i < index - 1; i++) {
            prev = prev.next;
        }

        Node n = prev.next;
        prev.next = n.next;

        return n;
    }

    @Test
    public void testAppend() {
        EX1_3_20 list = new EX1_3_20();
        list.append(1);
        list.append(2);
        list.append(3);

        assertEquals(Integer.valueOf(1), list.first.num);
        assertEquals(Integer.valueOf(2), list.first.next.num);
        assertEquals(Integer.valueOf(3), list.first.next.next.num);
    }

    @Test
    public void testDeleteFirst() {
        EX1_3_20 list = new EX1_3_20();
        list.append(1);
        list.append(2);
        list.append(3);
        assertEquals(Integer.valueOf(1), list.delete(0).num);
    }

    @Test
    public void testDeleteMiddle1() {
        EX1_3_20 list = new EX1_3_20();
        list.append(1);
        list.append(2);
        list.append(3);
        assertEquals(Integer.valueOf(2), list.delete(1).num);
    }

    @Test
    public void testDeleteMiddle2() {
        EX1_3_20 list = new EX1_3_20();
        list.append(1);
        list.append(2);
        list.append(3);
        list.append(4);
        assertEquals(Integer.valueOf(3), list.delete(2).num);
    }

    @Test
    public void testDeleteLast() {
        EX1_3_20 list = new EX1_3_20();
        list.append(1);
        list.append(2);
        list.append(3);
        assertEquals(Integer.valueOf(3), list.delete(2).num);
    }
}
