import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.StdRandom;
import edu.princeton.cs.algs4.Stopwatch;

public class SortCompare {
    public static double getElapsedTime(String alg, Comparable[] a) {
        Stopwatch timer = new Stopwatch();

        if (alg.equals("InsertionSorter")) InsertionSorter.sort(a);
        if (alg.equals("SelectionSorter")) SelectionSorter.sort(a);
        if (alg.equals("InsertionSorterV2")) InsertionSorterV2.sort(a);
        if (alg.equals("ShellSorter")) ShellSorter.sort(a);
        if (alg.equals("MergeSorter")) MergeSorter.sort(a);

        return timer.elapsedTime();
    }

    /**
     * 使用算法alg将T个长度为N的数组排序, 并返回总耗时
     *
     * @param algorithmName 排序算法的名称
     * @param arraySize     数组的长度
     * @param arrayNum      数组的个数
     * @return 总耗时
     */
    public static double timeRandomInput(String algorithmName, int arraySize, int arrayNum) {
        double totalTime = 0.0;
        Double[] sampleArray = new Double[arraySize];

        for (int n = 0; n < arrayNum; n++) {
            for (int i = 0; i < arraySize; i++) {
                sampleArray[i] = StdRandom.uniformDouble();
            }

            totalTime += getElapsedTime(algorithmName, sampleArray);
        }

        return totalTime;
    }

    public static void main(String[] args) {
        String alg1 = args[0];
        String alg2 = args[1];
        int N = Integer.parseInt(args[2]);
        int T = Integer.parseInt(args[3]);
        double t1 = timeRandomInput(alg1, N, T);
        double t2 = timeRandomInput(alg2, N, T);

        System.out.println(alg1 + " " + t1);
        System.out.println(alg2 + " " + t2);
        StdOut.printf("For %6d random Doubles: %s is %4.1f times faster than %s\n", N, alg1, t2 / t1, alg2);
    }
}
