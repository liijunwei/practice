+ "什么不应该可配置?"

+ 首先, 对于能够可靠的进行自动检测的东西, 就不需要提供配置开关; 警惕过早优化

+ 一个很好的经验法则是: 提高适应能力, 除非这样做会超过0.7秒的延迟(0.7秒是一个魔术)

+ 其次, 用户不应该看到优化开关; 让程序经济运行是**设计者**的任务, 不是用户的任务

+ KISS vs MICAHI(make it complicated and hide it)

+ 最后, 能用脚本包装器或者简单管道实现的任务, 就不要用配置开关实现; 能简单利用其他程序来完成的任务, 就不要增加本程序的复杂度
    + 我的 char_detector, 利用 `parallel` 加速, 而没有自己实现(机智)

# 10.5 命令行选项

有三种约定可以区分命令行选项和普通参数:

1. 原始的unix风格
    `-a`
    `-b`
    `-ab`
    `-c`
    ...

2. gnu风格
    `--config`
    `--debug`
    `--only`
    ...
    不支持组合
    选项和选项值之间 可以用 空格 或者 `=` 分隔
    良好的实践是: 至少对最常用的选项添加单字母等效选项


3. X toolkit 风格

## 10.5.1 从 -a 到 -z 的命令行选项(p243)

-I include

-k keep

-f file/force

-m message

-n number

-o output

-P protocol/port

-q quite

-s silent

-r/-R recurse/reverse

+ dang




