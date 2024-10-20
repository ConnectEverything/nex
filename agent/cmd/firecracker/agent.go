package main

import (
	"github.com/google/uuid"
	"github.com/synadia-io/nex/agent/agent"
)

// Firecracker agent implementation
type FirecrackerAgent struct {
	agent.Agent
}

func NewFirecrackerAgent() (*FirecrackerAgent, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &FirecrackerAgent{
		agent.Agent{
			ID: id.String(),
		},
	}, nil
}
