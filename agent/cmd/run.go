package cmd

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
	"github.com/synadia-io/nex/agent/agent"
)

const (
	runloopSleepInterval                = 250 * time.Millisecond
	runloopTickInterval                 = 2500 * time.Millisecond
	workloadExecutionSleepTimeoutMillis = 50
)

var (
	agt agent.AgentInterface

	cancelF context.CancelFunc
	closing uint32
	ctx     context.Context
	sigs    chan os.Signal

	// identifier string // FIXME-- either support an --id flag or remove

	natsURL    string
	nkeyPublic string
	nkeySeed   string
)

func init() {
	// TODO-- use something more robust instead of stdlib flag package
	// TODO-- add support for `up` and `preflight` commands

	// FIXME-- either support an --id flag or remove this `identifier` flag
	// it might be better to generate the id internally and
	// let the nex node configure how to refer to it by workload type...
	// flag.StringVar(&identifier, "id", "", "optional agent identifier; if not provided, a UUID will be generated and used")

	flag.StringVar(&natsURL, "nats-url", "", "url to use for the NATS connection")
	flag.StringVar(&nkeyPublic, "nkey-public", "", "public nkey to use for the NATS connection")
	flag.StringVar(&nkeySeed, "nkey-seed", "", "nkey seed to use for the NATS connection")
}

func Run(agt agent.AgentInterface) {
	flag.Parse()

	ctx, cancelF = context.WithCancel(context.Background())
	installSignalHandlers()

	nc, err := initNATSConnection()
	if err != nil {
		fmt.Printf("Failed to initialize NATS connection; %s\n", err.Error())
		shutdown()
		return
	}

	err = agt.Register(nc)
	if err != nil {
		fmt.Printf("Failed to register agent; %s\n", err.Error())
		shutdown()
		return
	}

	timer := time.NewTicker(runloopTickInterval)
	defer timer.Stop()

	for !shuttingDown() {
		select {
		case <-timer.C:
			// TODO
		case sig := <-sigs:
			fmt.Printf("Received signal: %s\n", sig)
			shutdown()
		case <-ctx.Done():
			shutdown()
		default:
			time.Sleep(runloopSleepInterval)
		}
	}
}

func initNATSConnection() (*nats.Conn, error) {
	if natsURL == "" {
		return nil, errors.New("--nats-url is required")
	}

	if nkeyPublic != "" {
		pair, err := nkeys.FromSeed([]byte(nkeySeed))
		if err != nil {
			return nil, err
		}

		return nats.Connect(natsURL,
			nats.DrainTimeout(time.Millisecond*5000),
			nats.Nkey(nkeyPublic,
				func(b []byte) ([]byte, error) {
					return pair.Sign(b)
				},
			),
		)
	}

	return nats.Connect(natsURL,
		nats.DrainTimeout(time.Millisecond*5000),
	)
}

func installSignalHandlers() {
	signal.Reset(syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
}

func shutdown() {
	if atomic.AddUint32(&closing, 1) == 1 {
		signal.Stop(sigs)
		close(sigs)

		if agt != nil {
			err := agt.Unregister()
			if err != nil {
				fmt.Printf("Failed to unregister agent; %s\n", err.Error())
			}

			// TODO-- agent.Shutdown()
			// HaltVM(nil)
		}
	}
}

func shuttingDown() bool {
	return (atomic.LoadUint32(&closing) > 0)
}
