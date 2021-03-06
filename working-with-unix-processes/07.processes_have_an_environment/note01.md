https://workingwithruby.com/wwup/env/

+ Environment, in this sense, refers to what’s known as `environment variables`

对进程来说, "环境"指的是环境变量

+ Environment variables are key-value pairs that hold data for a process.

环境变量是在进程里存储数据的键值对

+ Every process inherits environment variables from its parent.

每个进程从他的父进程里继承环境变量

+ **They are set by a parent process and inherited by its child processes.**

这些环境变量从父进程里继承而来, 并且会传递给它自己的子进程

+ Environment variables are per-process and are global to each process.

这些环境变量是每个进程的变量, 对每个进程都是全局变量

```ruby
MESSAGE='wing it' ruby -e "puts ENV['MESSAGE']"

# ENV in ruby is not a hash
ENV.class
=> Object
```

+ Environment variables are often used as a generic way to accept input into a command-line program

环境变量通常用作 接收命令行程序输入 的通用方法

+ Using environment variables is often less overhead than explicitly parsing command line options.

使用环境变量通常比显式解析命令行选项的开销更小

+ System Calls
    + man 3 setenv
    + man 3 getenv
    + man 7 environ

