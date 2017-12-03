# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure(2) do |config|
  config.vm.box = "bento/ubuntu-16.04"

  config.vm.network "forwarded_port", guest: 8081, host: 8081
  
  config.vm.synced_folder ".", "/home/vagrant/go/src/EasySurvey", create:true

  config.vm.provision "ansible" do |ansible|
    ansible.force_remote_user = false
    ansible.playbook = "ansible/main.yml"
  end
end
