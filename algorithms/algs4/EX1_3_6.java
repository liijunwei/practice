import edu.princeton.cs.algs4.Queue;
import edu.princeton.cs.algs4.Stack;
import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class EX1_3_6 {
    @Test
    public void testOutput() {
        Stack<Integer> stack = new Stack<>();
        Queue<Integer> q = new Queue<>();

        q.enqueue(1);
        q.enqueue(2);
        q.enqueue(3);
        q.enqueue(4);
        q.enqueue(5);

        assertEquals("1 2 3 4 5", q.toString().trim());

        while (!q.isEmpty()) {
            stack.push(q.dequeue());
        }

        while (!stack.isEmpty()) {
            q.enqueue(stack.pop());
        }

        assertEquals("5 4 3 2 1", q.toString().trim());
    }
}
