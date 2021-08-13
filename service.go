package nexus

import "net"

// Port is a network port that can be listened to for incoming traffic.
type Port struct {
	// UUID uniquely identifies this Port.
	UUID string

	// Number is the port number.
	Number int
}

// Endpoint describes a service connection endpoint, I.e. a socket.
type Endpoint struct {
	// Port is the port to be listened to for the service.
	Port Port

	// IP is the IP address on which the service port is being listened to.
	// A nil IP or an IP returning true for IP.IsUnspecified represents a
	// wildcard address, "0.0.0.0" or "::", I.e. a service listening on all
	// addresses.
	IP net.IP
}