# Skeleton Service Playbook

> Ansible Playbook for deploying the skeleton-service

## Requirements

* [Python](https://www.python.org/downloads/)
* [Vagrant](https://www.vagrantup.com/)
* [VirtualBox](https://www.virtualbox.org/wiki/Downloads)

## Installation

1. Install requirements listed above
2. ```./go precommit``` to validate the project is setup correctly

## Usage

    ./go install        Install dependencies
    ./go boot           Spin up a local virtual machine
    ./go install        Install all components to the local virtual machine
    ./go destroy        Destroy the local virtual machine
    ./go precommit      Run all validations before pushing code

## Acceptance Tests

This playbook has a suite of acceptance tests covering the main functionality of the playbook.

If you modify the playbook and want to verify your changes (and that you haven't broken anything else), run the tests.

To run all validations before pushing code, run ```./go precommit```

## Author

Jen Scobie (jenscobie@gmail.com)