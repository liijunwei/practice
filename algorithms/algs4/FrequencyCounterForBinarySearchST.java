import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;

import java.util.List;

public class FrequencyCounterForBinarySearchST {
    public static void main(String[] args) {
        int minlen = Integer.parseInt(args[0]);

        BinarySearchST<String, Integer> st = new BinarySearchST<>(2);

        while (!StdIn.isEmpty()) {
            String word = StdIn.readString();
            if (word.length() < minlen) {
                continue;
            }

            if (!st.contains(word)) {
                st.put(word, 1);
            } else {
                st.put(word, st.get(word) + 1);
            }
        }

        List<Integer> comparisonTimes = st.getComparisonTimes();
        VisualAccumulator va = new VisualAccumulator(comparisonTimes.size(), 6000);
        for (Integer comparisonTime : comparisonTimes) {
            va.addDataValue(comparisonTime);
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
