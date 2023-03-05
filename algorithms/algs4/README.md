[Algorithms 4th Edithion](https://algs4.cs.princeton.edu/home/)

https://github.com/kevin-wayne/algs4

+ TODO use Gradle

+ TODO resolve problems: https://lift.cs.princeton.edu/java/mac/

### Download data

https://introcs.cs.princeton.edu/java/data/

```bash
gem install nokogiri
make download_data
```

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

## 步骤(page 99)

在研究一个新的应用领域时, 我们将会按照以下步骤识别目标并使用数据抽象解决问题:

1. 定义API
2. 根据特定的应用场景开发用例代码
3. 描述一种数据结构(一组值的表示), 并在API所对应的抽象数据类型的实现中根据他定义类的实例变量
4. 描述算法(实现一组操作的方式), 并根据他实现类中的实例方法
5. 分析算法的性能特点


## 常见的错误

在编程领域中:

+ 最常见的错误就是过于关注程序的性能。你的首要任务应该是写出清晰正确的代码。仅仅为了提高运行速度而修改程序的事最好交给专家们来做。
+ 不成熟的优化是万恶之源
+ 另一个常见的错误是完全忽略了程序的性能

## Exercise References

https://github.com/reneargento/algorithms-sedgewick-wayne


