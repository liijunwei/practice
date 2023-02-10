package bag_queue_stack;

import org.junit.Test;

import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;

// practice 1.3.4
public class Parentheses {
    public static boolean foo(LinkedListStack<String> s1, String[] tokens) {
        for (String item : tokens) {
            if (item.equals("[") || item.equals("{") || item.equals("(")) {
                s1.push(item);
                continue;
            }

            String popStr = s1.pop();
            switch (item) {
                case "]":
                    if (popStr.equals("[")) {
                        continue;
                    } else {
                        return false;
                    }
                case "}":
                    if (popStr.equals("{")) {
                        continue;
                    } else {
                        return false;
                    }
                case ")":
                    if (popStr.equals("(")) {
                        continue;
                    } else {
                        return false;
                    }
                default:
                    return false;
            }
        }

        return s1.isEmpty();
    }

    @Test
    public void testBalancedParentheses1() {
        String input = "[()]{}{[()()]()}";
        String[] t = input.split("(?!^)");

        LinkedListStack<String> s = new LinkedListStack();
        assertTrue(foo(s, t));
    }

    @Test
    public void testBalancedParentheses2() {
        String input = "[(])";
        String[] t = input.split("(?!^)");
        LinkedListStack<String> s = new LinkedListStack();
        assertFalse(foo(s, t));
    }

    // @Ignore
    @Test
    public void testBalancedParentheses3() {
        String input = "[()]{}{[()()]()}[";
        String[] t = input.split("(?!^)");
        LinkedListStack<String> s = new LinkedListStack();
        assertFalse(foo(s, t));
    }
}
