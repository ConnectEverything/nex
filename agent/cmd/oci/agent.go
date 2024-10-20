package main

import (
	"github.com/google/uuid"
	"github.com/synadia-io/nex/agent/agent"
)

// OCI agent implementation
type OCIAgent struct {
	agent.Agent
}

func NewOCIAgent() (*OCIAgent, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &OCIAgent{
		agent.Agent{
			ID: id.String(),
		},
	}, nil
}
