package ThreadLocks;

public class HelloThread {
    public static void main(String[] args) throws InterruptedException {
        Thread t1 = new Thread(() -> System.out.println("hello from t1 thread"));

        t1.start();
        Thread.yield();
        System.out.println("hello from main thread");
        t1.join();
    }
}
