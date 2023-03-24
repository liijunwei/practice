import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdOut;
import org.junit.Test;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertTrue;

public class GraphTest {
    @Test
    public void testTinyG() {
        In in = new In("data/tinyG.txt");
        Graph G = new Graph(in);
        StdOut.println(G);

        assertTrue(true);
    }

    @Test
    public void testMediumG() {
        In in = new In("data/mediumG.txt");
        Graph G = new Graph(in);
        StdOut.println(G);

        assertTrue(true);
    }


    @Test
    public void testStaticMethods() {
        In in = new In("data/tinyG.txt");
        Graph G = new Graph(in);

        assertEquals(4, Graph.degree(G, 0));
        assertEquals(1, Graph.degree(G, 1));
        assertEquals(1, Graph.degree(G, 2));
        assertEquals(2, Graph.degree(G, 3));
        assertEquals(3, Graph.degree(G, 4));

        assertEquals(4, Graph.maxDegree(G));
        assertEquals(2.0, Graph.avgDegree(G), 0.001);

        assertEquals(0, Graph.numberOfSelfLoops(G));
    }
}
