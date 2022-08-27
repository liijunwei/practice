#!/bin/bash

# https://gist.github.com/puremourning/a41b4c6ac732091f63736e3ccb6d8d67

sudo yum install -y epel-release
sudo yum makecache

sudo yum groupinstall -y "Development tools"

sudo yum -y install https://packages.endpointdev.com/rhel/7/os/x86_64/endpoint-repo.x86_64.rpm
sudo yum install -y git
sudo yum install -y tig zsh


sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended && sudo chsh -s $(which zsh) vagrant