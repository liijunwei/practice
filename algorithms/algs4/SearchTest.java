import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdOut;
import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class SearchTest {
    private Search searcherFor(String inputFile, int source) {
        In in = new In(inputFile);
        Graph G = new Graph(in);

        Search search = new Search(G, source);

        for (int v = 0; v < G.V(); v++) {
            if (search.marked(v)) {
                StdOut.print(v + " ");
            }
        }
        StdOut.println();

        if (search.count() != G.V()) {
            StdOut.print("NOT ");
        }

        StdOut.println("connected");

        return search;
    }

    @Test
    public void testSearch() {
        assertEquals(7, searcherFor("data/tinyG.txt", 0).count());
        assertEquals(4, searcherFor("data/tinyG.txt", 9).count());
        assertEquals(2, searcherFor("data/tinyG.txt", 7).count());

    }
}
