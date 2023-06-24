package ThreadLocks.DiningPhilosophers;

import java.util.Random;

public class Philosopher extends Thread {
    final private String name;
    final private Chopstick left;
    final private Chopstick right;
    final private Random random;

    public Philosopher(String name, Chopstick left, Chopstick right) {
        this.name = name;
        this.left = left;
        this.right = right;
        this.random = new Random();
    }

    public void run() {
        try {
            while (true) {
                System.out.println(name + " thinking...");
                Thread.sleep(random.nextInt(1000));
                synchronized (left) {
                    synchronized (right) {
                        System.out.println(name + " eating with " + left.getId() + " and " + right.getId());
                        Thread.sleep(random.nextInt(1000));
                    }
                }
            }
        } catch (InterruptedException e) {
            System.out.println(e);
        }
    }

    public static void main(String[] args) throws InterruptedException {
        Philosopher[] philosophers = new Philosopher[5];
        Chopstick[] chopsticks = new Chopstick[5];

        for (int i = 0; i < 5; ++i)
            chopsticks[i] = new Chopstick(i);
        for (int i = 0; i < 5; ++i) {
            philosophers[i] = new Philosopher("P" + (i + 1), chopsticks[i], chopsticks[(i + 1) % 5]);
            philosophers[i].start();
        }
        for (int i = 0; i < 5; ++i)
            philosophers[i].join();
    }
}
