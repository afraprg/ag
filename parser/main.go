package parser

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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

	for n, flag := range cmd.flags() {
		if _, found := params[n]; found {
			for i, c2 := range cmd.CMDs {
				cmd.CMDs[i] = strings.Replace(c2, fmt.Sprintf("${%s}", n), params[n], 1)
			}
		} else if flag.Default != "" {
			for i, c2 := range cmd.CMDs {
				cmd.CMDs[i] = strings.Replace(c2, fmt.Sprintf("${%s}", n), flag.Default, 1)
			}
		} else if flag.Optional == false {
			log.Fatalf("Please enter %s", flag.Name)
		}
	}

	return cmd, err
}

func (c *Cmd) flags() map[string]Flag {
	result := make(map[string]Flag)
	for _, f := range c.Flags {
		result[f.Name] = f
	}

	return result
}
