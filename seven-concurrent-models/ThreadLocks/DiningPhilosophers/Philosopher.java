package ThreadLocks.DiningPhilosophers;

import java.util.Random;

public class Philosopher extends Thread {
    final private String name;
    final private Chopstick first;
    final private Chopstick second;
    final private Random random;

    public Philosopher(String name, Chopstick left, Chopstick right) {
        // 按照全局固定的顺序获取锁以避免死锁
        if (left.getId() < right.getId()) {
            this.first = left;
            this.second = right;
        } else {
            this.first = right;
            this.second = left;
        }

        this.name = name;
        this.random = new Random();
    }

    public void run() {
        try {
            while (true) {
                System.out.println(name + " thinking...");
                Thread.sleep(random.nextInt(1000));
                synchronized (first) {
                    synchronized (second) {
                        System.out.println(name + " eating with " + first.getId() + " and " + second.getId());
                        Thread.sleep(random.nextInt(1000));
                    }
                }
            }
        } catch (InterruptedException e) {
            System.out.println(e);
        }
    }

    public static void main(String[] args) throws InterruptedException {
        // increase this number will make dead lock easier to observe
//        int num = 5;
        int num = 50;
        Philosopher[] philosophers = new Philosopher[num];
        Chopstick[] chopsticks = new Chopstick[num];

        for (int i = 0; i < num; ++i)
            chopsticks[i] = new Chopstick(i);
        for (int i = 0; i < num; ++i) {
            philosophers[i] = new Philosopher("P" + (i + 1), chopsticks[i], chopsticks[(i + 1) % 5]);
            philosophers[i].start();
        }
        for (int i = 0; i < num; ++i)
            philosophers[i].join();
    }
}
