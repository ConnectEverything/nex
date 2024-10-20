package main

import (
	"github.com/google/uuid"
	"github.com/synadia-io/nex/agent/agent"
)

// FIXME?? I still dislike "DirectStart" as the name of this agent...
// Note that while this agent is packaged as a cmd/, it really needs
// to be imported and used directly by nex. The agent impl itself should
// be completely identical to the other agents that are installed by nex
// by way of the `nex preflight` command. The `NewDirectStartAgent` will
// be moved somewhere it can be imported and used by nex...

// DirectStartAgent agent implementation
type DirectStartAgent struct {
	agent.Agent
}

func NewDirectStartAgent() (*DirectStartAgent, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &DirectStartAgent{
		agent.Agent{
			ID: id.String(),
		},
	}, nil
}
