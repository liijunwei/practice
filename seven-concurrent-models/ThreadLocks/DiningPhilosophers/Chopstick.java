package ThreadLocks.DiningPhilosophers;

class Chopstick {
    private final int id;

    public Chopstick(int id) {
        this.id = id;
    }

    public int getId() {
        return id + 1;
    }
}
