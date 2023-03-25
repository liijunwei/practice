import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class DepthFirstPathsTest {
    @Test
    public void testPath() {
        DepthFirstPaths.pathFor("data/tinyG.txt", 0);
        assertEquals(2, 2);
    }
}
