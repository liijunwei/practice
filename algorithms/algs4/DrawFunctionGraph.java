// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/BinarySearch.java.html

import edu.princeton.cs.algs4.StdDraw;

public class DrawFunctionGraph {
    public static void main(String[] args) {
        int N = 100;
        StdDraw.setXscale(0, N);
        StdDraw.setYscale(0, N * N);
        StdDraw.setPenRadius(0.01);

        for(int i = 1; i <= N; i++){
            StdDraw.point(i, i);
            StdDraw.point(i, i * i);
            StdDraw.point(i, i * Math.log(i));
        }
    }
}
