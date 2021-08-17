package nexus

import "net"

// Nexus represents a network entity to which connections can be made.
type Nexus struct {
	// UUID uniquely identifies this Nexus.
	UUID string

	// Devs are the link-layer devices attached to this entity.
	Devices []*Device

	// Endpoints are the service endpoints being offered by this entity.
	Endpoints []*Endpoint
}

// HasIP returns true if the nexus has a device that is assigned the input IP.
func (n *Nexus) HasIP(ip net.IP) bool {
	for _, dev := range n.Devices {
		if dev.HasIP(ip) {
			return true
		}
	}
	return false
}
