import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;

import java.util.Arrays;

public class QuickUnion {
    private int[] id;
    private int count;

    public QuickUnion(int N) {
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
        while (p != id[p]) { // 相等时表示指向了自己
            p = id[p]; // 妙啊
        }

        return p;
    }

    // quick union

    /**
     * 连通p和q
     */
    public void union(int p, int q) {
        int pRoot = find(p);
        int qRoot = find(q);

        if (pRoot == qRoot) {
            return;
        }

        id[pRoot] = qRoot; // 指针

        count--;
    }

    public String toString() {
        return Arrays.toString(id);
    }

    public static void main(String[] args) {
        int N = StdIn.readInt();
        QuickUnion QuickUnion = new QuickUnion(N);

        while (!StdIn.isEmpty()) {
            int p = StdIn.readInt();
            int q = StdIn.readInt();

            if (!QuickUnion.connected(p, q)) {
                QuickUnion.union(p, q);
                StdOut.println(p + " " + q);
            }
        }

        StdOut.println("id: " + QuickUnion);
        StdOut.println(QuickUnion.count + " components");
    }
}
