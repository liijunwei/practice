# mini-agent1

一个极简的 ReAct（推理 → 行动）智能体，用 Clojure 实现，跑在 DeepSeek 的 Responses API 上，复用 Codex 的 API Key 和 URL。核心只有一个文件：HTTP 客户端 + 一个 `bash` 工具 + ReAct 循环 + 命令行入口。

## 配置

环境变量优先，缺省时自动读 `~/.codex/config.toml`：

| 项 | 环境变量 | 默认值 |
|----|----------|--------|
| API Key | `DEEPSEEK_API_KEY` | `~/.codex/config.toml` 的 `experimental_bearer_token` |
| Base URL | `DEEPSEEK_BASE_URL` | `https://api.deepseek.com/` |
| 模型 | `DEEPSEEK_MODEL` | `deepseek-flash` |

## 启动和使用

```bash
lein run -- "用 bash 查看 src 目录内容，并总结这个项目做什么"
lein run          # 不带参数进入交互式循环：每次 task> 输入一个任务，quit/exit/q 或 Ctrl-D 退出
```

## 工作原理

1. 把用户问题作为一条 `message` 发给 `/responses`。
2. 模型若返回 `function_call`，执行 `bash` 工具，把结果作为 `function_call_output` 追加进历史，回到第 1 步。
3. 模型若返回 `message` 文本，即最终答案，打印并结束。

每一步的工具调用和结果会实时打印到 stdout（`[tool] ...` / `[result] ...`），可以直接看到推理和行动过程。

## 调试

- 看中间过程：直接读 stdout 的 `[tool]` / `[result]`，不需要开关。
- 换 Key / 模型 / URL：`DEEPSEEK_API_KEY=sk-xxx DEEPSEEK_MODEL=... DEEPSEEK_BASE_URL=... lein run -- "..."`。
- 不联网调 ReAct 逻辑：

```clojure
(require '[mini-agent1.core :as a])
(a/react-loop (fn [_] [{:type "message" :content [{:type "output_text" :text "hi"}]}])
              (fn [_] "exit code: 0")
              "hi")
```

## 体验

```bash
lein run -- "列出当前目录的文件，说说你从这个项目里看到了什么"
lein run -- "运行 lein test 并报告结果"
```

## 验证

```bash
lein test   # 5 个单元测试：解析、工具执行、ReAct 循环、记忆
lein run -- "运行 lein test 并告诉我结果"   # 端到端验证
```

## 已知限制

- 只有一个 `bash` 工具，且不带沙箱：权限和你当前终端一致，能读写任何文件。
- 每轮任务最多 10 次工具调用，单条命令超时 60 秒。
- 多轮记忆只存在于单次 `lein run` 会话内（交互模式下跨轮保留），退出不落盘。
- 没有流式输出、文件编辑/搜索等工具。

## License

Copyright © 2026 mini-agent1 contributors. Distributed under the MIT License.
