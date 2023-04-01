import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;

import java.util.Arrays;

public class QuickFind {
    private int[] id;
    private int count;

    public QuickFind(int N) {
        this.count = N;
        this.id = new int[N];
        for (int i = 0; i < N; i++) {
            id[i] = i;
        }
    }

    public int count() {
        return count;
    }

    /**
     * 两个触点是否连通？
     * p和q是否存在于同一个分量中？
     */
    public boolean connected(int p, int q) {
        return find(p) == find(q);
    }

    /**
     * 找到p所在的分量(的标识符)
     */
    private int find(int p) {
        return id[p];
    }

    // quick union

    /**
     * 连通p和q
     */
    public void union(int p, int q) {
        int pID = find(p);
        int qID = find(q);

        if (pID == qID) {
            return;
        }

        for (int i = 0; i < id.length; i++) {
            // 数组内 和 pID 相等的值都要改为 qID，用这种方式表示连通
            // 最终数组内剩下几类数字，就表示最终有几个分量
            if (id[i] == pID) {
                id[i] = qID;
            }
        }

        count--;
    }

    public String toString() {
        return Arrays.toString(id);
    }

    public static void main(String[] args) {
        int N = StdIn.readInt();
        QuickFind quickFind = new QuickFind(N);

        while (!StdIn.isEmpty()) {
            int p = StdIn.readInt();
            int q = StdIn.readInt();

            if (!quickFind.connected(p, q)) {
                quickFind.union(p, q);
                StdOut.println(p + " " + q);
            }
        }

        StdOut.println("id: " + quickFind);
        StdOut.println(quickFind.count + " components");
    }
}
