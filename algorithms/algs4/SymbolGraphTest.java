import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;

public class SymbolGraphTest {
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
