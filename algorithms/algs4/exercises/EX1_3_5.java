package exercises;

import edu.princeton.cs.algs4.Stack;
import edu.princeton.cs.algs4.StdOut;
import org.junit.Test;

public class EX1_3_5 {
    @Test
    public void testOutput() {
        Stack<Integer> stack = new Stack<>();

        int N = 50;

        while (N > 0) {
            stack.push(N % 2);
            N = N / 2;
        }

        for (int d : stack) {
            StdOut.print(d);
        }
        StdOut.print();
        // 110010
    }
}
