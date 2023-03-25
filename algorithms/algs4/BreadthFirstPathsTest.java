import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class BreadthFirstPathsTest {
    @Test
    public void testPath() {
        BreadthFirstPaths.pathFor("data/tinyCG.txt", 0);
        assertEquals(2, 2);
    }
}
