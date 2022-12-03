package bag_queue_stack;

public class FixedCapacityStack<Item> {
    private Item[] stackEntries;
    private int size;

    public FixedCapacityStack(int cap) {
        stackEntries = (Item[]) new Object[cap]; // 强制类型转换
    }

    public void push(Item item) {
        if (size == stackEntries.length) {
            resize(2 * stackEntries.length);
        }

        stackEntries[size++] = item;
    }

    public Item pop() {
        Item item = stackEntries[--size];
        stackEntries[size] = null; // 避免对象游离(悬空)

        if (size > 0 && size == (stackEntries.length / 4)) {
            resize(stackEntries.length / 2);
        }

        return item;
    }

    public boolean isEmpty() {
        return size == 0;
    }

    public int size() {
        return size;
    }

    private void resize(int max) {
        Item[] temp = (Item[]) new Object[max];

        for (int i = 0; i < size; i++) {
            temp[i] = stackEntries[i];
        }

        stackEntries = temp;
    }
}
