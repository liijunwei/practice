package bag_queue_stack;

public class FixedCapacityStack<Item> {
    private Item[] stackEntries;
    private int size;

    public FixedCapacityStack(int cap) {
        stackEntries = (Item[]) new Object[cap]; // 强制类型转换
    }

    public void push(Item item) {
        stackEntries[size++] = item;
    }

    public Item pop() {
        return stackEntries[--size];
    }

    public boolean isEmpty() {
        return size == 0;
    }

    public int size() {
        return size;
    }
}
