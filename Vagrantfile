Vagrant.configure("2") do |config|
  config.vm.provider "virtualbox"
  config.vm.define "sandbox"
  # binding.pry
  if !`vagrant box list | grep junwei-centos-dev`.empty?
    
  else

  end

  config.vm.box = "junwei-centos-dev"
  config.vm.hostname = "practice"
  # config.vm.provision "shell", path: "provisioner.sh"
end
