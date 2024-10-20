package main

import (
	"github.com/google/uuid"
	"github.com/synadia-io/nex/agent/agent"
)

// Javascript agent implementation
type JavascriptAgent struct {
	agent.Agent
}

func NewJavascriptAgent() (*JavascriptAgent, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &JavascriptAgent{
		agent.Agent{
			ID: id.String(),
		},
	}, nil
}
