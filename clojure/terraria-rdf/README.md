# terraria-rdf / Terraria 晶塔小镇 NPC 分组优化器

基于 RDF 知识库 + SHACL 校验 + CP-SAT 整数规划（Quandary 封装 Google OR-Tools），
自动把 Terraria NPC 分成若干小镇（每组最多 3 人），在满足硬约束的前提下最小化总价格乘数。

## 核心思路

1. 用 RDF（Turtle）存储 NPC、生物群系和偏好知识，本体见 `resources/terraria-ontology.ttl`；
2. 用 SHACL 校验 RDF 数据，拦截冲突偏好和类型错误；
3. 用 CP-SAT 整数规划做最优分组；
4. 目标函数是最小化 log 得分总和（等价于最小化总价格乘数）。

## 硬约束

- 互相 `hatesNeighbor` 的 NPC 不能同镇；
- 一个小镇只能绑定单一生物群系；
- 小镇人数上限 = 3。

## 收益项

- 群系偏好、邻居偏好（单向/双向）加分；
- 公主同组自动给同组其他 NPC 加“喜爱”buff；
- 单人小镇独处折扣；
- 人数 > 2 时叠加拥挤惩罚。

## 项目结构（Leiningen）

```
terraria-rdf/
├── project.clj                    # 依赖、:main 入口、仓库镜像
├── resources/
│   ├── terraria-ontology.ttl      # 本体定义
│   ├── terraria-biomes.ttl        # 生物群系实体
│   ├── npc-facts.ttl              # NPC 偏好数据（已修正税收官）
│   └── terraria-shacl.ttl         # SHACL 校验规则
├── src/terraria_rdf/
│   ├── core.clj                   # -main 入口
│   └── solver.clj                 # RDF 加载 → SHACL 校验 → CP-SAT 求解
└── test/terraria_rdf/
    └── core_test.clj
```

## 运行前置

- Java 21 或更高版本（Jena 6 要求；本机已通过 `:java-cmd` 指向 Java 23，无需手动设置）；
- Leiningen（`lein`）。

## 运行

`project.clj` 已通过 `:java-cmd` 固定使用 Java 23，所以在项目目录下直接执行：

```bash
lein run
```

如果换机器，要么保证默认 `java` 是 21+，要么把 `project.clj` 里的 `:java-cmd` 改成该机器上
Java 21+ 的路径，或临时设置 `JAVA_HOME` 后再运行：

```bash
JAVA_HOME=<Java 21+ 的 HOME> lein run
```

程序会先做 SHACL 校验并打印单向偏好告警，然后启动 CP-SAT 求解并输出分组结果。

## 说明

- 依赖使用真实存在的版本：`io.github.mlimotte/quandary 0.9.1`、`com.dcj/clj-shacl 0.2.0`、
  `org.apache.jena/jena-arq 6.2.0`；另补了 Quandary 漏声明的 `org.clojure/tools.reader`。
- `project.clj` 里的 `:java-cmd` 指向本机 Java 23（Jena 6 需要 21+），这是机器相关的配置；
  迁移到其他机器时请按需修改。
- `project.clj` 里把 `central` 仓库指向了阿里云镜像（本机 IP 访问 Maven Central 被临时限流），
  若不需要可删除 `:repositories` 这一项。
