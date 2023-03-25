import union_find.WeightedQuickUnion;

public class Search {
    private final Graph g;
    private final int s;
    private final WeightedQuickUnion quickUnion;

    public Search(Graph g, int s) {
        this.g = g;
        this.s = s;

        this.quickUnion = new WeightedQuickUnion(g.V());
        for (int v = 0; v < g.V(); v++) {
            for (int w : g.adj(v)) {
                if (!quickUnion.connected(v, w)) {
                    quickUnion.union(v, w);
                }
            }
        }
    }

    public boolean marked(int v) {
        return quickUnion.connected(this.s, v);
    }

    public int count() {
        return quickUnion.countFor(s);
    }
}
