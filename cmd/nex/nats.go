package main

import (
	"strings"

	"github.com/nats-io/jsm.go/natscontext"
	"github.com/nats-io/nats.go"
)

func configureNatsConnection(cfg Globals) (*nats.Conn, error) {
	if cfg.Check {
		return nil, nil
	}

	opts := []nats.Option{
		nats.Name(cfg.NatsConnectionName),
	}

	// If a nats context is provided, it takes priority
	if !natscontext.IsKnown(cfg.NatsContext) {
		nc, err := natscontext.Connect(cfg.NatsContext, opts...)
		if err != nil {
			return nil, err
		}
		return nc, nil
	}

	if cfg.NatsTLSCert != "" && cfg.NatsTLSKey != "" {
		opts = append(opts, nats.ClientCert(cfg.NatsTLSCert, cfg.NatsTLSKey))
	}
	if cfg.NatsTLSCA != "" {
		opts = append(opts, nats.RootCAs(cfg.NatsTLSCA))
	}
	if cfg.NatsTLSFirst {
		opts = append(opts, nats.TLSHandshakeFirst())
	}

	switch {
	case cfg.NatsCredentialsFile != "": // Use credentials file
		opts = append(opts, nats.UserCredentials(cfg.NatsCredentialsFile))
	case cfg.NatsUserSeed != "" && cfg.NatsUserJWT != "": // Use seed + jwt
		opts = append(opts, nats.UserJWTAndSeed(cfg.NatsUserJWT, cfg.NatsUserSeed))
	case cfg.NatsUserNkey != "": // User nkey
		opts = append(opts, nats.Nkey(cfg.NatsUserNkey, nats.GetDefaultOptions().SignatureCB))
	case cfg.NatsUser != "" && cfg.NatsUserPassword != "": // Use user + password
		opts = append(opts, nats.UserInfo(cfg.NatsUser, cfg.NatsUserPassword))
	}

	nc, err := nats.Connect(strings.Join(cfg.NatsServers, ","), opts...)
	if err != nil {
		return nil, err
	}

	return nc, nil
}
