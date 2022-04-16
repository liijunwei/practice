#!/bin/bash

basedir=$(dirname -- "${BASH_SOURCE[0]}")
current_dir=$(pwd)

cd $basedir
rake
echo
find lib -name \*.rb | xargs flog
echo

cd $current_dir

