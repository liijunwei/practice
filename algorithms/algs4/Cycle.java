import edu.princeton.cs.algs4.In;

public class Cycle {
    private boolean[] marked;
    private boolean hasCycle;

    public Cycle(Graph g) {
        marked = new boolean[g.V()];

        for (int s = 0; s < g.V(); s++) {
            if (!marked[s]) {
                dfs(g, s, s);
            }
        }
    }

    private void dfs(Graph g, int v, int u) {
        marked[v] = true;

        for (int w : g.adj(v)) {
            if (!marked[w]) {
                dfs(g, w, v);
            } else if (w != u) {
                hasCycle = true;
            }
        }
    }

    public boolean hasCycle() {
        return hasCycle;
    }

    public static void main(String[] args) {
        String inputFile = args[0];
        In in = new In(inputFile);
        Graph g = new Graph(in);

        Cycle cycle = new Cycle(g);
        System.out.println(g);
        System.out.println("has cycle: " + cycle.hasCycle());
    }
}
