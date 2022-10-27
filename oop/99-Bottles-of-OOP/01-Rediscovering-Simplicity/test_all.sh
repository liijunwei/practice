#!/bin/bash

# practice && bash oop/99-Bottles-of-OOP/01-Rediscovering-Simplicity/test_all.sh

basedir=$(dirname -- "${BASH_SOURCE[0]}")
current_dir=$(pwd)

dirs=(
1.1.1.Incomprehensibly-Concise
1.1.2.Speculatively-General
1.1.3.Concretely-Abstract
1.1.4.Shameless-Green
solu_self
)

for dir in ${dirs[@]}; do
  cd "$basedir/$dir"
  echo "Run test for $dir"
  echo "============================"
  rake
  echo
  find lib -name \*.rb | xargs flog
  echo

  cd $current_dir
done

