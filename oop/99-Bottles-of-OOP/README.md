Recommended by [YimingChen](https://yiming.dev/blog/2017/12/08/book-review-99-bottles-of-oop/)

```bash
tig oop/99-Bottles-of-OOP/
```

# Setup

## install dependencies

```bash
gem list minitest
gem install minitest --version "~> 5.14"
```

## [Code Example](https://github.com/sandimetz/1st_99bottles_ruby)

The code is available in the 99bottles repository on GitHub, **which contains a branch for each chapter**.

```bash
git clone git@github.com:sandimetz/1st_99bottles_ruby.git
cd 1st_99bottles_ruby
tig --all
git bransh -a | cat
```

## [ABC metric tool](https://ruby.sadi.st/Flog.html)

```bash
gem install flog
find lib -name \*.rb | xargs flog
```

### what is ABC metric

https://stackoverflow.com/questions/30932732/what-is-meant-by-assignment-branch-condition-size-too-high-and-how-to-fix-it
https://en.wikipedia.org/wiki/ABC_Software_Metric#:~:text=The%20three%20components%20of%20the,Conditionals%3A%20Boolean%20or%20logic%20test.

The three components of the ABC score are defined as following:

+ Assignment: storage or transfer of data into a variable.
+ Branches: an explicit forward program branch out of scope.
+ Conditionals: Boolean or logic test.

# [Table of Contents](https://sandimetz.com/99bottles-sample-ruby)

> This book is about writing cost-effective, maintainable, and pleasing code.
>
> To get the most from the book, work the code samples as you read along.

