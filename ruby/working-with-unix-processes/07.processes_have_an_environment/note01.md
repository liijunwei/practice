https://workingwithruby.com/wwup/env/

+ Environment, in this sense, refers to what’s known as `environment variables`

+ Environment variables are key-value pairs that hold data for a process.

+ Every process inherits environment variables from its parent.

+ **They are set by a parent process and inherited by its child processes.**

+ Environment variables are per-process and are global to each process.

```ruby
MESSAGE='wing it' ruby -e "puts ENV['MESSAGE']"

# ENV in ruby is not a hash
ENV.class
=> Object
```

+ Environment variables are often used as a generic way to accept input into a command-line program

+ Using environment variables is often less overhead than explicitly parsing command line options.

+ System Calls
    + man 3 setenv
    + man 3 getenv
    + man 7 environ

