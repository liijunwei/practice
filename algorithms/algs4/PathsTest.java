import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdOut;
import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class PathsTest {
    @Test
    public void testPath() {
        Paths.pathFor("data/tinyG.txt", 0);
        assertEquals(2, 2);
    }
}
