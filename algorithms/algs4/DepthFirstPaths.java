import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.Stack;
import edu.princeton.cs.algs4.StdOut;

/**
 * 在G中找出所有起点为s的路径
 */
public class DepthFirstPaths {
    private boolean[] marked;
    private int[] edgeTo;
    private final int s;

    public DepthFirstPaths(Graph g, int s) {
        this.marked = new boolean[g.V()];
        this.edgeTo = new int[g.V()];
        this.s = s;
        dfs(g, s);
    }

    private void dfs(Graph g, int v) {
        marked[v] = true;

        for (int w : g.adj(v)) {
            if (!marked[w]) {
                edgeTo[w] = v;
                dfs(g, w);
            }
        }
    }

    public boolean hasPathTo(int v) {
        return marked[v];
    }

    public Iterable<Integer> pathTo(int v) {
        if (!hasPathTo(v)) {
            return null;
        }

        Stack<Integer> path = new Stack<>();
        for (int x = v; x != s; x = edgeTo[x]) {
            path.push(x);
        }

        path.push(s);

        return path;
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
