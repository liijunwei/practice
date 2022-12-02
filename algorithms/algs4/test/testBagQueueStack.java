package test;

import bag_queue_stack.Bag;
import org.junit.Test;

import java.util.Iterator;

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
}
