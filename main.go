package main

import (
	"ag/cmd"
	"ag/input"
	"ag/runner"
	"log"
)

func main() {
	cmd.Init()
	err := runner.Run(input.GetCmdName(), input.ParseArgs())
	if err != nil {
		log.Fatalln(err)
	}
}
