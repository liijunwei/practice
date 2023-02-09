package bag_queue_stack;

public class FixedCapacityStackOfStrings {
    private String[] stackEntries;
    private int size;
    private int cap;

    public FixedCapacityStackOfStrings(int cap) {
        this.cap = cap;
        stackEntries = new String[cap];
    }

    public void push(String item) {
        if (isFull()) {
            throw new RuntimeException("stack is full");
        }

        stackEntries[size++] = item;
    }

    public String pop() {
        return stackEntries[--size];
    }

    public boolean isEmpty() {
        return size == 0;
    }

    public boolean isFull() {
        return size >= cap;
    }

    public int size() {
        return size;
    }
}
