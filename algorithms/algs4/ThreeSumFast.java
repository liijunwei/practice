import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.Stopwatch;

import java.util.Arrays;

public class ThreeSumFast {
    public static int count(int[] a) {
        Arrays.sort(a);

        int N = a.length;

        int cnt = 0;
        for (int i = 0; i < N; i++) {
            for (int j = i + 1; j < N; j++) {
                int target = -(a[i] + a[j]);
                int k = BinarySearch.rank(target, a);
                if (k > j) {
                    cnt++;
                }
            }
        }

        return cnt;
    }

    public static void main(String[] args) {
        In in = new In(args[0]);
        int[] a = in.readAllInts();

        Stopwatch timer = new Stopwatch();
        int count = count(a);
        StdOut.println("elapsed time = " + timer.elapsedTime());
        StdOut.println(count);
    }
}
