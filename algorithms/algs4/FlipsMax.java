import edu.princeton.cs.algs4.Counter;
import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.StdRandom;

public class FlipsMax {
    public static Counter max(Counter x, Counter y) {
        return (x.tally() > y.tally()) ? x : y;
    }

    public static String toPercentage(Integer delta, Integer total) {
        return ((float) Math.abs(delta) / total) * 100 + "%";
    }

    public static void main(String[] args) {
        int total = Integer.parseInt(args[0]);

        Counter heads = new Counter("heads");
        Counter tails = new Counter("tails");

        for (int t = 0; t < total; t++) {
            if (StdRandom.bernoulli(0.5)) {
                heads.increment();
            } else {
                tails.increment();
            }
        }

        if (heads.tally() == tails.tally()) {
            StdOut.println("Tie");
        } else {
            StdOut.println(max(heads, tails) + " wins");
        }

        int d = heads.tally() - tails.tally();
        StdOut.println("diff: " + toPercentage(d, total));
    }
}
