import union_find.WeightedQuickUnion;

public class Search {
    private final Graph G;
    private final int s;
    private final WeightedQuickUnion quickUnion;

    public Search(Graph G, int s) {
        this.G = G;
        this.s = s;

        this.quickUnion = new WeightedQuickUnion(G.V());
        for (int v = 0; v < G.V(); v++) {
            for (int w : G.adj(v)) {
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
