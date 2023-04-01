import edu.princeton.cs.algs4.In;

public class CycleTest {
    public static void main(String[] args) {
        String inputFile = args[0];
        In in = new In(inputFile);
        Graph g = new Graph(in);

        Cycle cycle = new Cycle(g);
        System.out.println(g);
        System.out.println("has cycle: " + cycle.hasCycle());
    }
}
