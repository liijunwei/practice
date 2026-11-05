# mini-agent2

一个用 Clojure 实现的极简 ReAct coding agent。它沿用了
`mini-agent1` 的 DeepSeek Responses API、`bash` 工具和跨轮记忆，
并增加了 Ralph loop：每轮使用全新模型上下文，读取上一轮反馈，
反复修改项目，直到完成标记和验证命令同时通过。

## 环境

需要 Clojure、Leiningen、`bash`，以及一个 DeepSeek API Key。
Leiningen 会按 `project.clj` 自动安装 Clojure 与 `data.json` 依赖。

| 项 | 环境变量 | 默认值 |
|---|---|---|
| API Key | `MINI_AGENT2_API_KEY` 或 `DEEPSEEK_API_KEY` | Codex 配置的 bearer token |
| Base URL | `MINI_AGENT2_BASE_URL` 或 `DEEPSEEK_BASE_URL` | `https://api.deepseek.com/` |
| 模型 | `MINI_AGENT2_MODEL` 或 `DEEPSEEK_MODEL` | `deepseek-flash` |

## 使用

单次任务：

```bash
lein run -- "运行测试并总结当前项目"
```

不带参数进入交互模式，跨轮保留会话；`quit`、`exit`、`q` 或 Ctrl-D 退出：

```bash
lein run
```

运行 Ralph loop：

```bash
lein run -- ralph --verify "lein test" "实现 README 中描述的改动"
```

每轮使用干净上下文，最多 10 轮。达到轮数上限时以退出码 1 结束。
上一轮的答案和验证输出会反馈给下一轮，进度追加到
`.mini-agent2/ralph/progress.md`。

调整轮数和验证命令：

```bash
lein run -- ralph --max-iterations 5 --verify "lein test" "修复失败测试"
```

## 工作方式

1. 模型可以调用 `bash` 检查、编辑和测试当前工作目录。
2. ReAct 循环每轮最多执行 20 次工具调用，单条命令超时 60 秒。
3. Ralph 仅当答案含独立一行 `<promise>COMPLETE</promise>`，
   且验证命令退出码为 0 时结束。
4. 验证失败或模型未声明完成时，启动下一次全新上下文迭代。

## 验证

```bash
lein test
```

## 安全边界

`bash` 没有沙箱，权限与当前终端一致。Ralph loop 也会直接修改项目文件；
先在版本控制中提交或备份重要改动。
