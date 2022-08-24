Vagrant.configure("2") do |config|
  config.vm.provider "virtualbox"
  config.vm.define "sandbox"
  config.vm.box = "oppara/CentOS-7-dev"
  config.vm.hostname = "practice"
  config.vm.provision "shell", path: "provisioner.sh"
end
