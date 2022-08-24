Vagrant.configure("2") do |config|
  config.vm.provider "virtualbox"
  config.vm.define "sandbox"
  config.vm.hostname = "practice"

  use_local_box = ENV['SKIP_VAGRANT_BOX_CHECK'] || !`vagrant box list | grep junwei-centos-dev`.empty?
  if use_local_box
    config.vm.box = "junwei-centos-dev"
  else
    config.vm.box = "oppara/CentOS-7-dev"
    config.vm.provision "shell", path: "provisioner.sh"
  end
end
