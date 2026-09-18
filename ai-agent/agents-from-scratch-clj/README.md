# agents-from-scratch-clj

这是 [AI Agents from Scratch](https://github.com/pguso/agents-from-scratch) 的 Clojure 实现。
它不依赖 LangChain / ReAct 框架，按 12 个 lesson 逐步构建同一个 Agent。

参考仓库：<https://github.com/pguso/agents-from-scratch>

LLM 层调用 DeepSeek 的 OpenAI 兼容 chat API，默认模型为 `deepseek-v4-pro`。

## 核心流程

整个项目围绕一个不断变强的 Agent：

```text
用户输入
  -> LLM 调用（DeepSeek Chat Completions）
  -> JSON 抽取 / 校验 / 重试
  -> Agent 决策或工具调用
  -> 状态 / 记忆 / 计划 / 原子动作更新
  -> 返回结果
```

### 目录结构

```text
src/
├─ agent/       # Agent、状态、记忆、工具、计划、评估、遥测
├─ shared/      # LLM、JSON、prompt
├─ evals/       # 黄金测试集
├─ setup_check.clj
└─ complete_example.clj
```

### Lesson 对应关系

| Lesson | 能力             | 主要入口                     |
| ------ | ---------------- | ---------------------------- |
| 01     | 基础文本生成     | `agent/simple-generate`      |
| 02     | system prompt    | `agent/generate-with-role`   |
| 03     | 结构化 JSON 输出 | `agent/generate-structured`  |
| 04     | 决策路由         | `agent/decide`               |
| 05     | 工具调用         | `agent/request-tool`         |
| 06     | Agent loop       | `agent/run-loop`             |
| 07     | 记忆             | `agent/run-with-memory`      |
| 08     | 计划             | `agent/create-plan`          |
| 09     | 原子动作         | `agent/create-atomic-action` |
| 10     | AoT 依赖图       | `agent/create-aot-plan`      |
| 11     | Evals            | `agent.evals`                |
| 12     | Telemetry        | `agent.telemetry`            |

## 环境要求

- Java 17+
- [Leiningen](https://leiningen.org/)
- DeepSeek API Key
- 网络可访问 `api.deepseek.com`

## 配置

从环境变量读取 API Key：

```bash
export DEEPSEEK_API_KEY=你的key
```

可选：

```bash
export DEEPSEEK_MODEL=deepseek-v4-pro
```

默认模型是 `deepseek-v4-pro`。

## 执行

安装依赖：

```bash
lein deps
```

环境检查：

```bash
lein setup
```

运行全部 12 个 lesson：

```bash
lein run
# 或
lein example
```

只运行某一个模块：

```bash
lein run -m setup-check
lein run -m complete-example
```

## 测试

不调用真实 LLM 的单元测试：

```bash
lein test
```

当前测试覆盖：

- JSON 解析与抽取
- Agent state 与 memory
- calculator tool
- AoT 图执行

预期输出：

```text
Ran 4 tests containing 10 assertions.
0 failures, 0 errors.
```

## LLM 调用说明

`shared/llm.clj` 使用 `curl` 调用：

```text
POST https://api.deepseek.com/chat/completions
```

请求体：

```json
{
  "model": "deepseek-v4-pro",
  "messages": [{ "role": "user", "content": "..." }],
  "temperature": 0.0,
  "max_tokens": 4096,
  "stream": false,
  "stop": ["</s>", "\n\n", "User:", "Assistant:"]
}
```

`deepseek-v4-pro` 是 reasoning 模型，`max_tokens` 默认设为 `4096`，避免 reasoning tokens 占用后没有空间生成最终内容。
