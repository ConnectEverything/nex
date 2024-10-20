package agent

import (
	"errors"
	"fmt"
	"time"

	nats "github.com/nats-io/nats.go"
)

func (a *Agent) Register(nc *nats.Conn) error {
	if nc == nil {
		return errors.New("failed to register agent: nil NATS connection")
	}
	a.nc = nc

	err := a.requestHandshake()
	if err != nil {
		fmt.Printf("Failed to request handshake: %s\n", err)
		return err
	}

	deploySubject := fmt.Sprintf("agentint.%s.deploy", a.ID)
	sub, err := a.nc.Subscribe(deploySubject, a.handleDeploy)
	if err != nil {
		fmt.Printf("Failed to subscribe to agent deploy subject: %s\n", err)
		return err
	}
	a.subsz = append(a.subsz, sub)

	undeploySubject := fmt.Sprintf("agentint.%s.undeploy.*", a.ID)
	sub, err = a.nc.Subscribe(undeploySubject, a.handleUndeploy)
	if err != nil {
		fmt.Printf("Failed to subscribe to agent undeploy subject: %s\n", err)
		return err
	}
	a.subsz = append(a.subsz, sub)

	pingSubject := fmt.Sprintf("agentint.%s.ping", a.ID)
	sub, err = a.nc.Subscribe(pingSubject, a.handlePing)
	if err != nil {
		fmt.Printf("Failed to subscribe to ping subject: %s\n", err)
		return err
	}
	a.subsz = append(a.subsz, sub)

	return nil
}

func (a *Agent) Unregister() error {
	for _, sub := range a.subsz {
		if !sub.IsValid() {
			continue
		}

		_ = sub.Drain()
	}

	if a.nc != nil {
		_ = a.nc.Drain()
		for !a.nc.IsClosed() {
			time.Sleep(time.Millisecond * 25)
		}
	}

	return nil
}

func (a *Agent) handlePing(msg *nats.Msg) {
	_ = msg.Respond([]byte("OK"))
}

// Pull a deploy request off the wire, get the payload from the shared
// bucket, write it to tmp, initialize the execution provider per the
// request, and then validate and deploy a workload
func (a *Agent) handleDeploy(msg *nats.Msg) {
	// id, err := uuid.NewRandom()
	// if err != nil {
	// 	return nil, err
	// }
}

func (a *Agent) handleUndeploy(msg *nats.Msg) {
	// tokens := strings.Split(msg.Subject, ".")
	// workloadID := tokens[3]
}

func (a *Agent) requestHandshake() error {
	// a.submitLog("Requesting handshake from host", slog.LevelDebug)
	// msg := agentapi.HandshakeRequest{
	// 	ID:        a.md.VmID,
	// 	Message:   a.md.Message,
	// 	StartTime: a.started,
	//  Tenancy: a.tenancy,
	// }
	// raw, _ := json.Marshal(msg)

	// attempts := 0
	// for attempts < defaultAgentHandshakeAttempts-1 && !a.shuttingDown() {
	// 	attempts++

	// 	resp, err := a.nc.Request(fmt.Sprintf("hostint.%s.handshake", *a.md.VmID), raw, time.Millisecond*defaultAgentHandshakeTimeoutMillis)
	// 	if err != nil {
	// 		a.submitLog(fmt.Sprintf("Agent failed to request initial sync message: %s, attempt %d", err, attempts+1), slog.LevelError)
	// 		time.Sleep(time.Millisecond * 25)
	// 		continue
	// 	}

	// 	var handshakeResponse *agentapi.HandshakeResponse
	// 	err = json.Unmarshal(resp.Data, &handshakeResponse)
	// 	if err != nil {
	// 		a.submitLog(fmt.Sprintf("Failed to parse handshake response: %s", err), slog.LevelError)
	// 		time.Sleep(time.Millisecond * 25)
	// 		continue
	// 	}

	// 	a.submitLog(fmt.Sprintf("Agent is up after %d attempt(s)", attempts), slog.LevelInfo)
	return nil
}
