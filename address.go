package nexus

import (
	"net"
)

// Address represents an IP address on a Device.
type Address struct {
	// UUID uniquely identifies this Address.
	UUID string

	// DeviceUUID uniquely identifies the device
	// to which this address is attached.
	DeviceUUID string

	// IPNet is the IP network of this address.
	IPNet net.IPNet
}