#!/bin/bash

# https://gist.github.com/puremourning/a41b4c6ac732091f63736e3ccb6d8d67

sudo yum install -y epel-release
sudo yum makecache

sudo yum groupinstall -y "Development tools"

sudo yum install -y https://packages.endpointdev.com/rhel/7/os/x86_64/endpoint-repo.x86_64.rpm
sudo yum install -y git
sudo yum install -y tig
sudo yum install -y zsh
sudo yum install -y tldr
sudo yum install -y strace
sudo yum install -y bind-utils

sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended && sudo chsh -s $(which zsh) vagrant
