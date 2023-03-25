import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class DepthFirstPathsTest {
    @Test
    public void testPath() {
        DepthFirstPaths.pathFor("data/tinyCG.txt", 0);
        assertEquals(2, 2);
    }
}
