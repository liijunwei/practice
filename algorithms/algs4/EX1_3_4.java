import org.junit.Test;

import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;

// practice 1.3.4
public class EX1_3_4 {
    public static boolean parentheses(Stack<String> s, String[] tokens) {
        for (String item : tokens) {
            if (item.equals("[") || item.equals("{") || item.equals("(")) {
                s.push(item);
                continue;
            }

            if (s.isEmpty()) {
                return false;
            }

            String popStr = s.pop();
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

        return s.isEmpty();
    }

    @Test
    public void testBalancedParentheses1() {
        String input = "[()]{}{[()()]()}";
        String[] t = input.split("(?!^)");

        Stack<String> s = new Stack();
        assertTrue(parentheses(s, t));
    }

    @Test
    public void testBalancedParentheses2() {
        String input = "[(])";
        String[] t = input.split("(?!^)");
        Stack<String> s = new Stack();
        assertFalse(parentheses(s, t));
    }

    @Test
    public void testBalancedParentheses3() {
        String input = "[()]{}{[()()]()}[";
        String[] t = input.split("(?!^)");
        Stack<String> s = new Stack();
        assertFalse(parentheses(s, t));
    }

    @Test
    public void testBalancedParentheses_WithEmptyStack() {
        String input = "][";
        String[] t = input.split("(?!^)");
        Stack<String> s = new Stack();
        assertFalse(parentheses(s, t));
    }
}
