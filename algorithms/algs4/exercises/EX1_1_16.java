package exercises;

import org.junit.Test;

import static org.junit.Assert.*;

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

    // stackOverflow error
    public static String exR2(int n) {
        String s = exR2(n - 3) + n + exR2(n - 2) + n;
        if (n <= 0) return "";
        return s; // this line never get executed
    }

    @Test
    public void testExR2() {
        assertEquals("311361142246", exR2(6));
    }

    // damn 好神奇
    public static int mystery(int a, int b) {
        if (b == 0) return 0;
        if (b % 2 == 0) return mystery(a + a, b / 2);
        return mystery(a + a, b / 2) + a;
    }

    @Test
    public void testMystery() {
        assertEquals(50, mystery(2, 25));
        assertEquals(33, mystery(3, 11));
        assertEquals(12, mystery(2, 6));
        assertEquals(21, mystery(3, 7));
    }

    public static int mystery1(int a, int b) {
        if (b == 0) return 1;
        if (b % 2 == 0) return mystery1(a * a, b / 2);
        return mystery1(a * a, b / 2) * a;
    }

    @Test
    public void testMystery1() {
        assertEquals(33554432, mystery1(2, 25)); // 2^25
        assertEquals(177147, mystery1(3, 11));   // 3^11
    }
}
