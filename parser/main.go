package parser

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Cmd struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Flags       []Flag   `yaml:"flags"`
	CMDs        []string `yaml:"cmds"`
}

type Flag struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Optional    bool   `yaml:"optional"`
	Default     string `yaml:"default"`
}

func Parse(configFile string) (cmds []Cmd) {
	f, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalln("config file not found!")
	}
	err = yaml.Unmarshal(f, &cmds)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return
}
