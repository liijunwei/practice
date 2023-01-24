package exercises;

import org.junit.Test;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertTrue;

public class EX1_1_16 {
    public static String exR1(int n) {
        if (n <= 0) return "";
        // 注意是字符串操作，不是数字操作
        return exR1(n - 3) + n + exR1(n - 2) + n;
    }

    @Test
    public void testExR1() {
        assertEquals("311361142246", exR1(6));
    }

    @Test
    public void testTmp() {
        System.out.println("" + "" + 6);
        assertTrue(true);
    }
}
