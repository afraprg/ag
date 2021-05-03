# AG

![GO Version](https://img.shields.io/github/go-mod/go-version/afraprg/ag?style=for-the-badge) ![TAG](https://img.shields.io/github/v/tag/afraprg/ag?style=for-the-badge)

## Introduction

**ag** is a command line tool that similar to **Makefile**. with ag you can make an alias for group of commands with custom flags.

This tool is developed with golang.

## Installation

You can download the latest version on the [Release page](https://github.com/afraprg/ag/releases), after download you must move the binary file to **/usr/local/bin/ag** .

### Docker

In _<your_config_path>_ you must create a **.ag.yaml** file

    docker pull mostafahosseinime/ag
    docker run -v <your_config_path>:/root -it mostafahosseinime/ag:latest <your_command>

### Build from source

    go build .
    mv ag /usr/local/bin/ag

## Config
for the first step you have to create a config file, In the file below I have created an example of a config, you can create a file with your custom config.
```yaml
- name: glcp
  description: with this command you can create a goland project
  flags:
    - name: pn
      description: project name
      optional: false
      default: project 1
    - name: pt
      description: location
      optional: true
      default: "/Users/mostafahosseini/go/src/"
  cmds:
    - cd ${pt}
    - mkdir ${pn}
    - cd ${pn}
    - git init
    - touch main.go
    - git add .
    - goland .
```
after create the config file you must run the following command:

       $ ag init your_config.yaml

## Run

after run init command with that config you can create a new golang project and finaly open in goland with this command:

       $ ag glcp --pn=example_project --pt=/users/mostafahosseini/go/src/example_project

also, I created [an example file](https://github.com/afraprg/ag/blob/main/.docker/.ag.yaml "an example file") in this repository.


