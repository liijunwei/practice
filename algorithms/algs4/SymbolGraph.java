import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.ST;
import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;

public class SymbolGraph {
    private ST<String, Integer> st; // 符号名 -> 索引
    private String[] keys;          // 索引 -> 符号名
    private Graph graph;

    public SymbolGraph(String stream, String sp) {
        st = new ST<>();
        In in = new In(stream);

        while (in.hasNextLine()) {
            String[] a = in.readLine().split(sp);

            for (String s : a) {
                if (!st.contains(s)) {
                    st.put(s, st.size());
                }
            }
        }

        keys = new String[st.size()];

        for (String name : st.keys()) {
            keys[st.get(name)] = name;
        }

        graph = new Graph(st.size());
        in = new In(stream);

        while (in.hasNextLine()) {
            String[] a = in.readLine().split(sp);

            int v = st.get(a[0]);

            for (int i = 1; i < a.length; i++) {
                graph.addEdge(v, st.get(a[i]));
            }
        }
    }

    public boolean contains(String s) {
        return st.contains(s);
    }

    public int index(String s) {
        return st.get(s);
    }

    public String name(int v) {
        return keys[v];
    }

    public Graph graph() {
        return graph;
    }

    public static void main(String[] args) {
        String inputFile = args[0];
        String delim = args[1];

        SymbolGraph sg = new SymbolGraph(inputFile, delim);
        Graph g = sg.graph();

        while (StdIn.hasNextLine()) {
            String source = StdIn.readLine();
            for (int w : g.adj(sg.index(source))) {
                StdOut.println("  " + sg.name(w));
            }
        }
    }
}
