### Install clojure and open repl

```bash
# https://clojure.org/guides/install_clojure

brew install clojure/tools/clojure
# disable netskope and run `clj` to download dependencies and open repl
clj
```

### clojure之道---分离标识与状态

+ clojure提供了一种方法, 混合了两数式编程和可变状态-—这种方法平衡了两者优点, 成为并发编程的利器

+ persistent data structure
    + "持久数据结构"里的"持久"指的不是将数据持久化到磁盘或者数据库中，而是指数据结构被修改时总是保留其之前的版本，这样可以为代码提供一致的数据视角
    + 持久数据结构被修改时看上去就像创建了一个完整的副本
    + [持久数据结构的实现选择了"共享结构"](https://github.com/matthiasn/talk-transcripts/blob/master/Hickey_Rich/PersistentDataStructure.md)
```clojure
user=> (def mapv1 {:name "paul" :age 45})
#'user/mapv1
user=> (def mapv2 (assoc mapv1 :gender :male))
#'user/mapv2
user=> mapv1
{:name "paul", :age 45}
user=> mapv2
{:name "paul", :age 45, :gender :male}
```

+ 我们这里说的“持久” 并不是指将数据持久化到磁盘或者保存到数据库中，而是指数据结构 被修改时总是保留其之前的版本，这样可以为代码提供 一致的数据视角。来用REPL看一个简单 的 例 子:

+ 原子变量就是具有原子性的变量, 非常类似于2.3节中介绍的原子变量(事实上Clojure的原 子变量就是在java. util.concurrent.atomic的基础上建立的)




