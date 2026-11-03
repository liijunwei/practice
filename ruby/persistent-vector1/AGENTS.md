# AGENTS.md

从零复刻 Clojure 持久化数据结构的 Ruby 学习项目（Hypirion「Understanding Clojure's Persistent Vectors」）。一条铁律贯穿所有结构：改动返回新值、旧值不变。实现手段是「路径复制」——只复制根到改动点的链，其余节点共享；节点数组在共享前一律 `freeze`。

## 文件职责

- `persistent_vector.rb` — `PersistentVector`：带 tail 的正式版（系列第 3 篇）。
- `persistent_vector_basic.rb` — `PersistentVectorNoTail`：无 tail 基线，**只**给 `bench_tail.rb` 当对照。没有测试；别改它的行为、别加功能，否则基准失效。
- `persistent_map.rb` — `PersistentMap`：不可变哈希表，同一套 trie，路由依据换成 `key.hash` 低 30 位。
- `bench_tail.rb` — tail vs 无 tail 基准。
- `test_*.rb` — Minitest。`test_persistent_vector_properties.rb` 是手写 PBT：固定种子 + Random 生成，差分对照 Ruby Array/Hash，无 shrink。

## 改动约定

- 递归路径复制的每个 base-case 分支都显式 `return`；漏掉会 SystemStackError（`pop_leaf` 被 PBT 抓到过一次）。
- 结构 bug 几乎都藏在 32 的幂边界（向量 33/65/1057/32801）。改结构时先让属性测试在这些点上红过，再修到绿。
- 手写 PBT 没有 shrink，失败靠 seed/round 复现；保留固定种子，别删。

## 环境与坑

- 本机 Ruby 不支持 endless method（`def x = ...`），统一写 `def x ... end`。
- 无 Gemfile/Rake。跑全部测试：`for f in test_*.rb; do ruby "$f" || break; done`
- `PersistentMap` 不可变：没有 `[]=`，写操作用 `assoc`/`dissoc` 并重新绑定变量。`==` 只认 map 与 Hash，且不对称——`pm == hash` 真、`hash == pm` 假。
