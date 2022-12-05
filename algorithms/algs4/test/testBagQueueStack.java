package test;

import bag_queue_stack.*;
import edu.princeton.cs.algs4.StdOut;
import org.junit.Test;

import static org.junit.Assert.*;

public class testBagQueueStack {
    @Test
    public void testBagWithString() {
        Bag<String> bag = new Bag<>();

        assertEquals(0, bag.size());
        assertTrue(bag.isEmpty());

        bag.add("你好，明天");
        bag.add("你好，生活");
        assertEquals(2, bag.size());
        assertFalse(bag.isEmpty());
    }

    @Test
    public void testBagWithDate() {
        Bag<Integer> bag = new Bag<>();

        assertEquals(0, bag.size());
        assertTrue(bag.isEmpty());

        bag.add(1988);
        bag.add(2);
        assertEquals(2, bag.size());
        assertFalse(bag.isEmpty());
    }

    @Test
    public void testIterateThroughBag() {
        Bag<String> bag = new Bag<>();

        bag.add("你好醫生，你能不能把我殺了");
        bag.add("我的胸口，好像被他們堵住了");
        bag.add("掏了半天什麼都沒有，但能感覺心跳加快");
        bag.add("突發奇想去一個地方，到了之後又想回來");

        for (String s : bag) {
            System.out.println(s);
        }

        assertEquals(4, bag.size());
    }

    @Test
    public void testStats() {
        Bag<Double> bag = new Bag<>();

        bag.add(100.0);
        bag.add(99.0);
        bag.add(101.0);
        bag.add(120.0);
        bag.add(98.0);
        bag.add(107.0);
        bag.add(109.0);
        bag.add(81.0);
        bag.add(101.0);
        bag.add(90.0);

        double sum = 0;
        for (Double d : bag) {
            sum += d;
        }

        int N = bag.size();
        double mean = sum / N;

        sum = 0;
        for (Double d : bag) {
            sum += (d - mean) * (d - mean);
        }

        double std = Math.sqrt(sum / N);

        assertEquals("mean", 100.6, mean, 0.1);
        assertEquals("std", 9.9, std, 0.1);
    }

    @Test
    public void testFixedCapacityStackOfStrings() {
        String tobe = "to be or not to - be - - that - - - is";
        String[] tokens = tobe.split(" ");

        FixedCapacityStackOfStrings s = new FixedCapacityStackOfStrings(10);
        assertEquals(0, s.size());

        for (String item : tokens) {
            if (!item.equals("-")) {
                s.push(item);
            } else if (!s.isEmpty()) {
                StdOut.print(s.pop() + " ");
            }
        }

        assertEquals(2, s.size());
        assertEquals("is", s.pop());
        assertEquals("to", s.pop());
        assertEquals(0, s.size());
    }

    @Test
    public void testFixedCapacityStack() {
        String tobe = "to be or not to - be - - that - - - is";
        String[] tokens = tobe.split(" ");

        FixedCapacityStack<String> s = new FixedCapacityStack(2);
        assertEquals(0, s.size());

        for (String item : tokens) {
            if (!item.equals("-")) {
                s.push(item);
            } else if (!s.isEmpty()) {
                StdOut.print(s.pop() + " ");
            }
        }

        assertEquals(2, s.size());
        assertEquals("is", s.pop());
        assertEquals("to", s.pop());
        assertEquals(0, s.size());
    }

    @Test
    public void testResizingArrayStack() {
        ResizingArrayStack<Integer> s = new ResizingArrayStack(2);
        assertEquals(0, s.size());
        assertTrue(s.isEmpty());

        s.push(1);
        s.push(2);
        s.push(3);
        s.push(4);
        s.push(5);
        assertEquals(5, s.size());

        StringBuilder str = new StringBuilder();
        for (Integer i : s) {
            str.append(i);
        }
        assertEquals("54321", str.toString());

    }

    @Test
    public void testLinkedListStack() {
        String tobe = "to be or not to - be - - that - - - is";
        String[] tokens = tobe.split(" ");

        LinkedListStack<String> s = new LinkedListStack();
        assertEquals(0, s.size());
        assertTrue(s.isEmpty());

        for (String item : tokens) {
            if (!item.equals("-")) {
                s.push(item);
            } else if (!s.isEmpty()) {
                StdOut.print(s.pop() + " ");
            }
        }

        assertEquals(2, s.size());
        assertEquals("is", s.pop());
        assertEquals("to", s.pop());
        assertEquals(0, s.size());

    }
}
