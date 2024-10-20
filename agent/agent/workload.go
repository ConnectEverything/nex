package agent

import "github.com/google/uuid"

func (a *Agent) Deploy(req *DeployRequest) (*Workload, error) {
	return nil, nil
}

// Undeploy a workload
func (a *Agent) Undeploy(id uuid.UUID) error {
	return nil
}
