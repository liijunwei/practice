import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class TestRedBlackBST {
    @Test
    public void testPut() {
        RedBlackBST<String, String> tree = new RedBlackBST();
        tree.put("description", "red black bst demo");
        tree.put("name", "foo");
        tree.put("age", "11");

        assertEquals(3, tree.size());
    }
}
