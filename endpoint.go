package nexus

import (
	"fmt"
	"net"
)

// Endpoint describes a service connection endpoint, I.e. a socket.
type Endpoint struct {
	// UUID uniquely identifies this Endpoint.
	UUID string

	// Nexus is a reference to the Nexus that offers this endpoint.
	Nexus *Nexus

	// Port is the network port to be listened to by the service.
	Port int

	// IP is the IP address on which the service port is being listened to.
	// A nil IP or an IP returning true for IP.IsUnspecified represents a
	// wildcard address, "0.0.0.0" or "::", I.e. a service listening on all
	// addresses.
	IP net.IP
}

// WithEndpointUUID returns a function that sets the input UUID on an Endpoint.
func WithEndpointUUID(uuid string) func(*Endpoint) {
	return func(e *Endpoint) {
		e.UUID = uuid
	}
}

// OnNexus returns a function that sets the
// input Nexus reference on an Endpoint.
func OnNexus(nex *Nexus) func(*Endpoint) {
	return func(e *Endpoint) {
		e.Nexus = nex
	}
}

// NewEndpoint returns a new Endpoint based on the input arguments.
func NewEndpoint(port int, options ...func(endpoint *Endpoint)) (*Endpoint, error) {
	if port < 0 || port > 65535 {
		return nil, fmt.Errorf("port %d is outside the allowable range 0-65535", port)
	}

	ep := &Endpoint{
		Port: port,
	}

	for _, opt := range options {
		opt(ep)
	}

	return ep, nil
}
