public class Puzzle {
    static boolean answerReady = false;
    static int answer = 0;

    static Thread t1 = new Thread(() -> {
        answer = 42;
        answerReady = true;
    });

    static Thread t2 = new Thread(() -> {
        if (answerReady) {
            System.out.println("The meaning of life is: " + answer);
        } else {
            System.out.println("I don't know the answer");
        }
    });

    // there is a chance that the output is: "The meaning of life is: 0"
    public static void main(String[] args) throws InterruptedException {
        t1.start();
        t2.start();
        t1.join();
        t2.join();
    }
}
