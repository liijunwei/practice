// https://www.geeksforgeeks.org/producer-consumer-solution-using-threads-java/

import java.util.concurrent.locks.ReentrantLock;

public class demo {
    private static class Counter {
        private int count;
        private final int n;
        private final ReentrantLock mutex;

        public Counter(int n) {
            this.count = 0;
            this.mutex = new ReentrantLock();
            this.n = n;
        }

        public void increment() {
            this.count++;
        }

        public void decrement() {
            this.count--;
        }

        public boolean can_produce() {
            return this.count < this.n;
        }

        public boolean can_consume() {
            return this.count > 0;
        }
    }

    public static void t_produce(Counter counter) {
        while (true) {
            counter.mutex.lock();

            if (counter.can_produce()) {
                counter.increment();
                System.out.print("(");
            }

            counter.mutex.unlock();
        }
    }

    public static void t_consume(Counter counter) {
        while (true) {
            counter.mutex.lock();

            if (counter.can_consume()) {
                counter.decrement();
                System.out.print(")");
            }

            counter.mutex.unlock();
        }
    }

    public static void main(String[] args) throws InterruptedException {
        int n = Integer.parseInt(args[0]);
        Counter c = new Counter(n);
        int t = Integer.parseInt(args[1]);

        for (int i = 0; i < t; i++) {
            Thread tp = new Thread(() -> t_produce(c));
            Thread tc = new Thread(() -> t_consume(c));

            tp.start();
            tc.start();

            tp.join();
            tc.join();
        }
    }
}
