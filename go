#!/bin/bash

set -e

#VBoxManage -v >/dev/null 2>&1 || { echo >&2 "VirtualBox is required. Please install the latest version."; exit 1; }
#vagrant -v >/dev/null 2>&1 || { echo >&2 "Vagrant is required. Please install the latest version."; exit 1; }
#python -V >/dev/null 2>&1 || { echo >&2 "Python is required. Please install the latest version."; exit 1; }
#
#[[ $(vagrant plugin list) == *vagrant-vbguest* ]] || { vagrant plugin install vagrant-vbguest; }
#
#REQUIRED_RUBY=2.1.2
#
#(rbenv versions | grep $REQUIRED_RUBY) || rbenv install $REQUIRED_RUBY
#rbenv local $REQUIRED_RUBY
#(rbenv exec gem list | grep bundler) || rbenv exec gem install bundler
#bundle --path=vendor/bundle --quiet

function helptext {
    echo "Usage: ./go <command>"
    echo ""
    echo "Available commands are:"
    echo "    boot         Spin up a local virtual machine"
    echo "    destroy      Destroy the local virtual machine"
    echo "    execute      Provision EC2 servers"
    echo "    precommit    Run all validations before pushing code"
    echo "    provision    Provision the local virtual machine"
    echo "    setup        Install project dependencies"
    echo "    spec         Run acceptance tests against the local virtual machine"
}

function boot {
    vagrant up --no-provision
}

function install_roles {
    ansible-galaxy install --role-file=Galaxyfile --roles-path=roles --force
}

function provision {
    install_roles
    boot
    vagrant provision
}

function destroy {
    vagrant destroy -f
}

function execute() {
    install_roles
    is_known_environment $1
    ansible-playbook site.yml -i hosts -l $1 --private-key ~/.ssh/ec2.pem -u ec2-user
}

function is_known_environment() {
    case $1 in
        development|qa|production)
        ;;
        *)  echo "$1 not in list of known environments (development|qa|production)"
            exit 1
        ;;
    esac
}

function setup {
    which pip >/dev/null 2>&1 || sudo easy_install pip
    ansible --version  >/dev/null 2>&1 || sudo pip install ansible
}

function spec {
    bundle exec rake spec
}

function precommit {
    setup
    provision
    spec
}

[[ $@ ]] || { helptext; exit 1; }

case "$1" in
    boot) boot
    ;;
    destroy) destroy
    ;;
    execute) execute $2
    ;;
    help) helptext
    ;;
    precommit) precommit
    ;;
    provision) provision
    ;;
    setup) setup
    ;;
    spec) spec
    ;;
esac