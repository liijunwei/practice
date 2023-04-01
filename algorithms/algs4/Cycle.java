/**
 * 检查图里是否存在环
 */
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
        System.out.printf("%2d is marked, dfs(g, %2d, %2d)\n", v, v, u);

        // 理不清楚...
        for (int w : g.adj(v)) {
            if (!marked[w]) {
                dfs(g, w, v);
            } else if (w != u) {
//                System.out.printf("has cycle when: v = %2d w = %2d u = %2d\n", v, w, u);
                hasCycle = true;
            }
        }
    }

    public boolean hasCycle() {
        return hasCycle;
    }
}
