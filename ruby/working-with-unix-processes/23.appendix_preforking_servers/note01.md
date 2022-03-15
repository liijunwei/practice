# [Appendix: Preforking Servers](https://workingwithruby.com/wwup/prefork/)

+ At the core of all these projects is the preforking model. There are a few things about preforking that make it special, here are 3:
    1. Efficient use of memory.
    2. Efficient load balancing.
    3. Efficient sysadminning.

+ Preforking uses memory more efficiently than does spawning multiple unrelated processes.

+ The truth is that it does take some time to fork a process (it’s not instantaneous瞬间) and that there is some memory overhead for each child process. These values are negligible可以忽略不计的 compared to the overhead of booting many Mongrels. Preforking wins.

+ **Keep in mind that the benefits of copy-on-write are forfeited if you're running MRI. To reap these benefits you need to be using REE.**
    + Q: what is REE?
    + A: [Ruby Enterprise Edition](https://rvm.io/interpreters/ree)
    + REE builds on top of MRI Rubies, versions 1.8.X and later, to deliver an enhanced interpreter with many performance and memory optimizations, including common patch-sets such as MBARI.



