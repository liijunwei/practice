import edu.princeton.cs.algs4.StdDraw;
import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;

public class FrequencyCounter {
    private static class VisualAccumulator {
        public VisualAccumulator() {
            StdDraw.setXscale(0, 200);
            StdDraw.setYscale(0, 10);
            StdDraw.setPenRadius(0.005);
        }

        public void addDataValue(int x, int y, double mean) {
            StdDraw.setPenColor(StdDraw.DARK_GRAY);
            StdDraw.point(x, y);

            StdDraw.setPenColor(StdDraw.RED);
            StdDraw.point(x, mean);
        }
    }

    public static void main(String[] args) {
        int minlen = Integer.parseInt(args[0]);
        VisualAccumulator va = new VisualAccumulator();

        SequentialSearchST<String, Integer> st = new SequentialSearchST<>();

        int totalPutCount = 0;
        int index = 0;

        while (!StdIn.isEmpty()) {
            index++;
            int putCount = 0;

            String word = StdIn.readString();
            if (word.length() < minlen) {
                va.addDataValue(index, putCount, (double) totalPutCount / index);

                continue;
            }

            if (!st.contains(word)) {
                st.put(word, 1);
            } else {
                st.put(word, st.get(word) + 1);
            }

            putCount++;
            totalPutCount += putCount;
            va.addDataValue(index, putCount, (double) totalPutCount / index);
        }

        String max = " ";
        st.put(max, 0);

        for (String word : st.keys()) {
            if (st.get(word) > st.get(max)) {
                max = word;
            }
        }

        StdOut.println(max + " " + st.get(max));
    }
}
