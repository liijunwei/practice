import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdOut;

/**
 * 在G中找出所有起点为s的路径
 */
public class DepthFirstPaths {
    public DepthFirstPaths(Graph g, int s) {

    }

    public boolean hasPathTo(int v) {
        return false;
    }

    public Iterable<Integer> pathTo(int v) {
        return null;
    }

    public static DepthFirstPaths pathFor(String inputFile, int source) {
        In in = new In(inputFile);
        Graph g = new Graph(in);

        DepthFirstPaths search = new DepthFirstPaths(g, source);

        for (int v = 0; v < g.V(); v++) {
            StdOut.print(source + " to " + v + ": ");

            if (search.hasPathTo(v)) {
                for (int x : search.pathTo(v)) {
                    if (x == source) {
                        StdOut.print(x);
                    } else {
                        StdOut.print("-" + x);
                    }
                }
            }

            StdOut.println();
        }
        StdOut.println();

        return search;
    }

    public static void main(String[] args) {
        String inputFile = args[0];
        int s = Integer.parseInt(args[1]);
        DepthFirstPaths.pathFor(inputFile, s);
    }
}
