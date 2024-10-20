package main

import (
	"github.com/synadia-io/nex/agent/cmd"
)

func main() {
	agent, err := NewOCIAgent()
	if err != nil {
		panic(err)
	}

	cmd.Run(agent)
}
