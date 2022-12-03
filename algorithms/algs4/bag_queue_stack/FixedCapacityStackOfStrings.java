package bag_queue_stack;

public class FixedCapacityStackOfStrings {
    private String[] stackEntries;
    private int size;

    public FixedCapacityStackOfStrings(int cap) {
        stackEntries = new String[cap];
    }

    public void push(String item) {
        stackEntries[size++] = item;
    }

    public String pop() {
        return stackEntries[--size];
    }

    public boolean isEmpty() {
        return size == 0;
    }

    public int size() {
        return size;
    }
}
