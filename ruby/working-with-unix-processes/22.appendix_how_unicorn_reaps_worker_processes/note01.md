# [Appendix: How Unicorn Reaps Worker Processes](https://workingwithruby.com/wwup/unicorn/)

+ Unicorn web server

+ What’s the big deal?
    + Unicorn is a web server that attempts to push as much responsibility onto the kernel as it can.
    + It uses lots of Unix Programming. The codebase is chock full of Unix Programming techniques.

+ Before we dive into the code I’d like to provide a bit of context about how Unicorn works.
+ At a very high level Unicorn is a pre-forking web server.



