package agent

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/jwt/v2"
	nats "github.com/nats-io/nats.go"
)

// Agents implement this interface
// FIXME-- this should live in a public nex package but is here temporarily during dev
type AgentInterface interface {
	// Deploy a workload
	Deploy(req *DeployRequest) (*Workload, error)

	// Undeploy a workload
	Undeploy(id uuid.UUID) error

	// Run preflight
	Preflight() error

	// Register the agent with a nex node using the given NATS connection
	Register(nc *nats.Conn) error

	// Unregister the agent
	Unregister() error

	// TODO-- add interface methods for interrogating running workloads
}

// FIXME-- this should reference the public nex package type but is here temporarily during dev
type DeployRequest struct {
	// FIXME-- do not duplicate Workload struct
}

// FIXME-- this should reference the public nex package type but is here temporarily during dev
type Workload struct {
	Argv          []string          `json:"argv,omitempty"`
	DecodedClaims jwt.GenericClaims `json:"-"`
	Description   *string           `json:"description"`
	Environment   map[string]string `json:"environment"`
	Essential     *bool             `json:"essential,omitempty"`
	FunctionID    *string           `json:"function_id,omitempty"`
	Hash          string            `json:"hash,omitempty"`
	ID            *string           `json:"id"`
	Location      *url.URL          `json:"location"`
	Name          string            `json:"workload_name,omitempty"`
	Namespace     string            `json:"namespace,omitempty"`
	RetriedAt     *time.Time        `json:"retried_at,omitempty"`
	RetryCount    *uint             `json:"retry_count,omitempty"`
	TotalBytes    int64             `json:"total_bytes,omitempty"`

	// HostServicesConfig *controlapi.NatsJwtConnectionInfo `json:"host_services_config,omitempty"`
	TriggerSubjects []string `json:"trigger_subjects"`
	Type            string   `json:"type,omitempty"`

	// Stderr      io.Writer `json:"-"`
	// Stdout      io.Writer `json:"-"`
	// TmpFilename *string   `json:"-"`

	EncryptedEnvironment *string `json:"-"`
	JsDomain             *string `json:"-"`
	SenderPublicKey      *string `json:"-"`
	WorkloadJWT          *string `json:"-"`

	Errors []error `json:"errors,omitempty"`
}
