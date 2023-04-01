import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class JunitTest {
    @Test
    public void testSetup() {
        String str = "I am done with Junit setup";
        assertEquals("I am done with Junit setup", str);
    }

    @Test
    public void testStringOperation() {
        String a = "now is ";
        String b = "the time ";
        String c = "to";

        assertEquals(a.length(), 7);
        assertEquals(a.charAt(4), 'i');
        assertEquals(a.concat(c), "now is to");
        assertEquals(a.indexOf("is"), 4);
        assertEquals(a.substring(2, 5), "w i");
        assertEquals(a.split(" ")[0], "now");
        assertEquals(a.split(" ")[1], "is");
        assertEquals(b.equals(c), false);
    }
}
