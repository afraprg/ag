package cmd

import (
	"ag/parser"
	"io"
	"log"
	"os"
)

func Init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) == 1 {
		os.Exit(0)
	}

	if os.Args[1] == "init" && len(os.Args) == 2 {
		os.Exit(0)
	} else if os.Args[1] == "init" && os.Args[2] != "" {
		copyConfig(os.Args[2], dirname+"/.ag.yaml")
		os.Exit(0)
	}

	parser.ConfigFile = dirname + "/.ag.yaml"
}

func copyConfig(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
