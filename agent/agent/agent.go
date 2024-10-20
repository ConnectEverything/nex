package agent

import (
	nats "github.com/nats-io/nats.go"
)

// Base agent
type Agent struct {
	ID string

	nc    *nats.Conn
	subsz []*nats.Subscription

	workloads map[string]*Workload
}
