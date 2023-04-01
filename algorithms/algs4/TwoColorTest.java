import edu.princeton.cs.algs4.In;

public class TwoColorTest {
    public static void main(String[] args) {
        String inputFile = args[0];
        In in = new In(inputFile);
        Graph g = new Graph(in);

        TwoColor color = new TwoColor(g);
        System.out.println(g);
        System.out.println("is bipartite: " + color.isBipartite());
    }
}
