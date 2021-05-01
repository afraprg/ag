package input

import (
	"os"
	"strings"
)

func GetCmdName() string {
	return os.Args[1]
}

func ParseArgs() map[string]string {
	if len(os.Args[2:]) < 1 {
		return nil
	}
	output := make(map[string]string)
	for _, arg := range os.Args[2:] {
		av := strings.Split(strings.Replace(arg, "--", "", 1), "=")
		output[av[0]] = av[1]
	}

	return output
}
