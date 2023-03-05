import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;

public class FrequencyCounter {
    public static void main(String[] args) {
        int minlen = Integer.parseInt(args[0]);
        VisualAccumulator va = new VisualAccumulator(10000, 10000);

        SequentialSearchST<String, Integer> st = new SequentialSearchST<>();

        while (!StdIn.isEmpty()) {
            String word = StdIn.readString();
            if (word.length() < minlen) {
                continue;
            }

            int count;
            if (!st.contains(word)) {
                count = st.put(word, 1);
            } else {
                count = st.put(word, st.get(word) + 1);
            }

            va.addDataValue(count);
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
