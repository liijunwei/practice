import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdOut;
import org.junit.Test;

import static org.junit.Assert.assertTrue;

public class SearchTest {
    @Test
    public void testSearch0() {
        In in = new In("data/tinyG.txt");
        Graph G = new Graph(in);
        int s = 0;

        Search search = new Search(G, s);

        for(int v =0; v < G.V();v++){
            if(search.marked(v)){
                StdOut.print(v+" ");
            }
        }
        StdOut.println();

        if(search.count() !=G.V()){
            StdOut.print("NOT ");
        }

        StdOut.println("connected");

        assertTrue(true);
    }

    @Test
    public void testSearch9() {
        In in = new In("data/tinyG.txt");
        Graph G = new Graph(in);
        int s = 9;

        Search search = new Search(G, s);

        for(int v =0; v < G.V();v++){
            if(search.marked(v)){
                StdOut.print(v+" ");
            }
        }
        StdOut.println();

        if(search.count() !=G.V()){
            StdOut.print("NOT ");
        }

        StdOut.println("connected");

        assertTrue(true);
    }
}
