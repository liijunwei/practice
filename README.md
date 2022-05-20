# Practice

> [Work consistently, but don’t overwork yourself.](http://blog.thefirehoseproject.com/posts/learn-to-code-and-be-self-reliant/)

This repo means nothing but my programming practices and book reading notes.

```bash
# tig dirpath to view directory history
# e.g.
tig mysql/mysql-crash-cource
tig mysql/imooc-learn-117/
tig cpp/
tig clang/the-c-programming-language/
tig working-with-unix-processes/
```

## [auto commit](https://www.markusdosch.com/2020/07/git-auto-commit-push-every-couple-of-minutes/)

```bash
watch -n 60 "git pull && (git ls-files --modified --others --exclude-standard | grep . > /dev/null) && { git add . ; git commit -m 'Commit automatically by watch.' ; git push; }"
```
