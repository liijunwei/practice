[Algorithms 4th Edithion](https://algs4.cs.princeton.edu/home/)

+ TODO use Gradle

+ TODO resolve problems: https://lift.cs.princeton.edu/java/mac/

### 链表

+ 链表是一种递归的数据结构, 它或者为空, 或者是含有泛型元素的节点和另一个节点的指针
+ 链表表示的是一列元素

+ 节点记录

```java
private class Node {
  Item item;
  Node next;
}
```

+ 操作
    + 构建链表
    + 在 表头 插入节点(prepend)
    + 在 表头 删除节点(delete_first)
    + 在 表尾 插入节点(append)
    + 在 表尾 删除节点(delete_last)
    + 在 其他位置 插入节点(insert_at)
    + 在 其他位置 删除节点(delete_at)
    + 删除 指定节点
    + 在 指定节点前 插入节点(insert_before)
    + 在 指定节点后 插入节点(insert_after)
    + 遍历链表


+ 用法
    + 实现(下压)栈Stack
    + 实现(先进先出)队列Queue
