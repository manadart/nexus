package nexus

import (
	"net"
)

// Device represents a physical or virtual network device that is addressable
// at the link-layer via a hardware address, such as MAC or InfiniBand GUID.
type Device struct {
	// UUID uniquely identifies this device.
	UUID string

	// Nexus is a reference to the Nexus to which this Device is attached.
	Nexus *Nexus

	// HWAddr is the link-layer address of the device.
	HWAddr string

	// NetNSID identifies the network namespace for this device.
	// See https://en.wikipedia.org/wiki/Linux_namespaces#Network_(net).
	NetNSID string

	// Addresses are the IP addresses attached to this device.
	Addresses []*Address
}

// HasIP returns true if the device is assigned the input IP.
func (d *Device) HasIP(ip net.IP) bool {
	for _, addr := range d.Addresses {
		if addr.IP.Equal(ip) {
			return true
		}
	}
	return false
}