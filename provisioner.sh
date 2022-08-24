#!/bin/bash

# https://gist.github.com/puremourning/a41b4c6ac732091f63736e3ccb6d8d67

sudo yum groupinstall -y "Development tools"
sudo yum install -y git
sudo yum install -y zsh

sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended


