package nexus

import (
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