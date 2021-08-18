package nexus

import (
	"fmt"
	"net"
)

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

// AddEndpoint adds the input Endpoint to the Nexus unless it is already
// present or is assigned to another Nexus, in which cases an error results.
// TODO: Add tests.
func (n *Nexus) AddEndpoint(ep *Endpoint) error {
	if ep.Nexus != nil && ep.Nexus.UUID != n.UUID {
		return fmt.Errorf("endpoint is already associated with another nexus %s", ep.Nexus.UUID)
	}

	for _, nEP := range n.Endpoints {
		if nEP.UUID == ep.UUID {
			return fmt.Errorf("endpoint %s is already associated with this nexus", ep.UUID)
		}
		if nEP.Port == ep.Port {
			if ep.IP == nil || nEP.IP == nil || ep.IP.Equal(nEP.IP) {
				return fmt.Errorf("nexus %s already has an endpoint for port %d", n.UUID, ep.Port)
			}
		}
	}

	ep.Nexus = n
	n.Endpoints = append(n.Endpoints, ep)
	return nil
}
