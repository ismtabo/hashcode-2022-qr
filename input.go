package main

import (
	"io/ioutil"
	"os"
)

func getInput() string {
	var input []byte
	if len(os.Args) < 2 {
		var err error
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
	} else {
		var err error
		input, err = ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	return string(input)
}

type Problem struct {
}

func parseInput(input string) (problem Problem) {
	return
}
