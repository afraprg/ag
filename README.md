# AG

//todo

## Introduction

**ag** is a command line tool that similar to **Makefile**. with ag you can make an alias for groups of command with custom flags.

This tool is developed with golang.

## Installation

//todo

### Docker

    docker pull mostafahosseinime/ag

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

after run init command with that config you can create a new golang project and finaly open in goland with this command:

    ag glcp --pn=example_project --pt=/users/mostafahosseini/go/src/example_project

also, I created [an example file](https://github.com/afraprg/ag/blob/main/.docker/.ag.yml "an example file") in this repository.


