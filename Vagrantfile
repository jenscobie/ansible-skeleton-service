# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "chef/centos-6.6"

  config.vm.provision :ansible do |ansible|
    ansible.playbook = "site.yml"
  end
end
