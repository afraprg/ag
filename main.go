package main

import (
	"ga/cmd"
	"ga/input"
	"ga/runner"
	"log"
)

func main() {
	cmd.Init()
	err := runner.Run(input.GetCmdName(), input.ParseArgs())
	if err != nil {
		log.Fatalln(err)
	}
}
