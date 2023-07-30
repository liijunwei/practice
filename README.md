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

## [auto commit](https://github.com/liijunwei/omz-git/blob/4c85c101e11ef9fd660f4378c8f51ca60b26fbdc/aliases.sh#L19)

```bash
alias git-setup-auto-commit="watch -n 120 \"git pull && git add . && git commit -m 'Commit automatically by watch.'\""
```

## vm for c programs

```bash
vagrant up
vagrant ssh
cd /vagrant && ls
```

## example of using linux env

```bash
git clone https://github.com/liijunwei/practice
cd practice && vagrant up && vagrant ssh
cd /vagrant/csapp/ch08-exceptional-control-flow && make ecf_fork

# or

vagrant ssh-config
# copy the config to your ~/.ssh/config file
# use vscode "Remote - SSH" plugin
Remote-SSH connect to sandbox
```
