package main

import (
	"github.com/synadia-io/nex/agent/cmd"
)

func main() {
	agent, err := NewJavascriptAgent()
	if err != nil {
		panic(err)
	}

	cmd.Run(agent)
}
