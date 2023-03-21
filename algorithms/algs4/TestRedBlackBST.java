import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class TestRedBlackBST {
    @Test
    public void testPut1() {
        RedBlackBST<String, String> tree = new RedBlackBST();
        tree.put("description", "red black bst demo");
        tree.put("name", "foo");
        tree.put("age", "11");

        assertEquals(3, tree.size());
        tree.print();
    }

    @Test
    public void testPut2() {
        RedBlackBST<String, String> tree = new RedBlackBST();
        tree.put("S", "foo1");
        tree.put("E", "foo2");
        tree.put("A", "foo3");
        tree.put("R", "foo4");
        tree.put("C", "foo5");
        tree.put("H", "foo6");
        tree.put("X", "foo7");
        tree.put("M", "foo8");
        tree.put("P", "foo9");
        tree.put("L", "foo10");

        assertEquals(10, tree.size());
        tree.print();
    }
    @Test
    public void testGet() {
        RedBlackBST<String, String> tree = new RedBlackBST();
        tree.put("A", "foo1");
        tree.put("C", "foo2");
        tree.put("E", "foo3");
        tree.put("H", "foo4");
        tree.put("L", "foo5");
        tree.put("M", "foo6");
        tree.put("P", "foo7");
        tree.put("R", "foo8");
        tree.put("S", "foo9");
        tree.put("X", "foo10");

        assertEquals(10, tree.size());
        assertEquals("foo10", tree.get("X"));
        assertEquals("foo5", tree.get("L"));
    }
}
