# 开发阶段 REPL、AOT 与 uberjar 学习笔记

日期：2026-10-14

## 背景

在 Calva / VS Code 中执行 Jack-in 时，输出里先出现了这样一段：

```text
This is not a "real" terminal.
You can type into the terminal, but there is no process that will handle your input.

Connecting using "Leiningen" project type.
Starting Jack-in: (cd /Users/me/codeout/practice/clojure/fwpd; lein update-in ... -- with-profile +uberjar repl :headless)
```

这里整理了当时提出的问题、回答，以及开发阶段启动 REPL 的建议。

## 问题 1：AOT compile 是什么意思

AOT 是 Ahead-Of-Time Compilation，也就是“提前编译”。

Clojure 通常是在运行时才把 `.clj` 代码交给 JVM 编译成字节码。AOT 编译则是在启动前、构建时就把 `.clj` 源文件编译成 `.class` 字节码文件。

典型目的：

- 加快程序启动速度
- 生成可以部署的 JAR 包
- 配合 `direct-linking` 等优化手段

## 问题 2：`:aot :all` 是什么意思

`:aot` 是 Leiningen 的编译配置。

它后面可以指定命名空间，也可以写成 `:all`。` :all` 表示把 `src/` 下的所有命名空间都做 AOT 编译。

本项目在 [`project.clj`](../project.clj) 中的配置是：

```clojure
:profiles {:uberjar {:aot :all
                     :jvm-opts ["-Dclojure.compiler.direct-linking=true"]}}
```

## 问题 3：uberjar profile 是什么意思

`profile` 是 Leiningen 的一组配置切片，用于在不同场景下启用不同的依赖、JVM 参数、编译选项等。

`uberjar` profile 是项目里专门用于打“uberjar”的配置。uberjar 是指把项目代码和所有依赖打包成一个可以直接运行的独立 JAR 文件。

激活方式通常是：

```bash
lein uberjar
```

或者在执行其他任务时临时合并该 profile：

```bash
lein with-profile +uberjar <task>
```

## 问题 4：uberjar 和“必须编译通过才能启动 REPL”是什么关系

Calva 的 Jack-in 命令末尾是：

```bash
with-profile +uberjar repl :headless
```

执行顺序如下：

1. `with-profile +uberjar` 先把 `uberjar` profile 合并进来。
2. 该 profile 里有 `:aot :all`。
3. 因此 Leiningen 在真正启动 nREPL 之前，会先编译所有命名空间。
4. 如果任何命名空间编译失败，整个 `repl` 任务就会中止，nREPL 不会启动。

所以：

```text
with-profile +uberjar
        ↓
激活 :aot :all
        ↓
编译 fwpd.core 等所有命名空间
        ↓
编译失败 → repl 任务中止 → 看不到 nREPL server
编译成功 → 才会启动 nREPL server
```

普通的 `lein repl` 不会一开始就做全量 AOT 编译，所以通常 REPL 可以先起来；等真正加载有问题的命名空间时才会报错。

## 本次 Jack-in 失败的根因

失败不是那句 “This is not a real terminal”，那只是 Calva 输出窗口的说明。

真正原因是：项目使用了 `with-profile +uberjar`，而 `uberjar` profile 里有 `:aot :all`。AOT 编译 `fwpd.core` 时，[`src/fwpd/core.clj`](../src/fwpd/core.clj) 顶层表达式执行出错。

当时的错误写法是：

```clojure
(mapv #(clojure.string/split % #"\n") (slurp filename))
```

`slurp` 返回字符串，`mapv` 会把字符串拆成一个一个字符传给函数；而 `clojure.string/split` 需要的是 `CharSequence`，不是 `Character`，所以抛出：

```text
class java.lang.Character cannot be cast to class java.lang.CharSequence
```

建议改为：

```clojure
(str/split-lines (slurp filename))
```

但更根本的建议是：不要把这些探索性表达式直接放在命名空间顶层，详见下一节。

## 开发阶段怎么启动 REPL

### 总体原则

开发阶段不需要 AOT 编译，也不应该用 uberjar profile 启动 REPL。

`+uberjar` 是“打包发布”场景用的；日常开发更适合普通的项目 REPL。

### 推荐方式一：终端直接启动 REPL

在项目根目录运行：

```bash
lein repl
```

如果需要给 Calva 连接，用 headless 模式：

```bash
lein repl :headless
```

启动后会用 nREPL 监听一个端口，然后在 Calva 中执行：

```text
Calva: Connect to a running REPL
```

并输入终端显示的 `host:port`。

### 推荐方式二：让 Calva 使用开发 profile

如果继续使用 Calva 的 Jack-in，最好为 Leiningen 配置一个不含 `+uberjar` 的 connect sequence。

核心思路是：让 Jack-in 使用普通 `repl`，不要合并 `uberjar` profile。这样启动 REPL 时就不会先做 `:aot :all`。

### 推荐方式三：让命名空间顶层保持干净

不管用哪种方式启动 REPL，都建议避免在 `ns` 顶层放“直接执行的探索代码”。

可以把探索性表达式放进 `(comment ...)`：

```clojure
(comment
  (str/split-lines (slurp filename))

  (mapv #(clojure.string/split % #",")
        (str/split-lines (slurp filename)))

  (->> (slurp filename)
       str/split-lines
       (map #(str/split % #","))
       (map row->map-1)))
```

这样：

- 文件加载或 AOT 编译时不会执行这些探索表达式
- 不会因为临时实验代码导致 REPL 启动失败
- 需要查看结果时，手动把光标放到表达式上，在 Calva 中单独求值即可

## 小结

- 开发阶段推荐 `lein repl`，不要用 `with-profile +uberjar`。
- `uberjar` profile 用于打包，不用于日常 REPL。
- `:aot :all` 会在启动前编译所有命名空间，因此任何顶层错误都会阻止 nREPL 启动。
- 探索性代码应放在 `(comment ...)` 中，而不是直接放在命名空间顶层。
