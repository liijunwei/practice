# [Appendix: Preforking Servers](https://workingwithruby.com/wwup/prefork/)

+ At the core of all these projects is the preforking model. There are a few things about preforking that make it special, here are 3:
    1. Efficient use of memory.
    2. Efficient load balancing.
    3. Efficient sysadminning.

+ Preforking uses memory more efficiently than does spawning multiple unrelated processes.

