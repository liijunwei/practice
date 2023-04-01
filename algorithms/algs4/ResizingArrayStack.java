import java.util.Iterator;

public class ResizingArrayStack<Item> implements Iterable<Item> {
    private Item[] stackEntries;
    private int size;

    public ResizingArrayStack(int cap) {
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

        if (size >= 0) {
            System.arraycopy(stackEntries, 0, temp, 0, size);
        }

        stackEntries = temp;
    }

    @Override
    public Iterator<Item> iterator() {
        return new ReverseArrayIterator();
    }

    private class ReverseArrayIterator implements Iterator<Item> {
        // 支持LIFO的迭代(遍历)
        private int i = size;

        public boolean hasNext() {
            return i > 0;
        }

        public Item next() {
            return stackEntries[--i];
        }

        public void remove() {
        }
    }
}
