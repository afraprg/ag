package parser

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

var ConfigFile string

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

func parse() ([]Cmd, error) {
	f, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		return nil, err
	}
	var CMDs []Cmd
	err = yaml.Unmarshal(f, &CMDs)
	if err != nil {
		return nil, err
	}

	return CMDs, nil
}

func Find(name string, params map[string]string) (Cmd, error) {
	CMDs, err := parse()
	if err != nil {
		return Cmd{}, err
	}
	var cmd Cmd
	for _, c := range CMDs {
		if c.Name == name {
			cmd = c
			break
		}
	}

	for pk, pv := range params {
		for i, c2 := range cmd.CMDs {
			cmd.CMDs[i] = strings.Replace(c2, pk, pv, 1)
		}
	}

	return cmd, err
}
