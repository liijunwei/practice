import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;

import java.util.List;

public class FrequencyCounter {
    public static void main(String[] args) {
        int minLength = Integer.parseInt(args[0]);

        SequentialSearchST<String, Integer> st = new SequentialSearchST<>();

        while (!StdIn.isEmpty()) {
            String word = StdIn.readString();
            if (word.length() < minLength) {
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
