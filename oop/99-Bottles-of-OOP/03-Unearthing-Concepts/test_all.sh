#!/bin/bash

# practice && bash oop/99-Bottles-of-OOP/03-Unearthing-Concepts/test_all.sh

basedir=$(dirname -- "${BASH_SOURCE[0]}")
current_dir=$(pwd)

cd $basedir
rake
echo
cd $current_dir

