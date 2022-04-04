#!/bin/bash

basedir=$(dirname -- "${BASH_SOURCE[0]}")
current_dir=$(pwd)

cd $basedir
rake
echo
cd $current_dir

