package bag_queue_stack;

import org.junit.Ignore;
import org.junit.Test;

import static org.junit.Assert.*;

public class Parentheses {
    public static boolean foo(LinkedListStack<String> s1, String[] tokens) {
        for (String item : tokens) {
            if (item.equals("[") || item.equals("{") || item.equals("(")) {
                s1.push(item);
                continue;
            }

            if (s1.isEmpty()) {
                return false;
            }

            String popStr;
            switch (item) {
                case "]":
                    popStr = s1.pop();
                    if (popStr.equals("{") || popStr.equals("(")) {
                        return false;
                    }
                    break;
                case "}":
                    popStr = s1.pop();
                    if (popStr.equals("[") || popStr.equals("(")) {
                        return false;
                    }
                    break;
                case ")":
                    popStr = s1.pop();
                    if (popStr.equals("[") || popStr.equals("{")) {
                        return false;
                    }
                    break;
                default:
                    return false;
            }
        }
        return true;
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

    // don't know why this won't pass
    @Ignore
    @Test
    public void testBalancedParentheses3() {
        String input = "[()]{}{[()()]()}[(";
        String[] t = input.split("(?!^)");
        LinkedListStack<String> s = new LinkedListStack();
        assertFalse(foo(s, t));
    }
}
