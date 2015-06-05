# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "ubuntu/trusty64"

  config.vm.network "private_network", ip: "192.168.50.4"

  config.vm.provision :ansible do |ansible|
    ansible.limit = 'all'
    ansible.playbook = "provision.yml"
    ansible.inventory_path = "hosts"
    ansible.sudo = true
  end
end
