package nexus

import (
	"net"
	"strings"
)

// Address represents an IP address on a Device.
type Address struct {
	// UUID uniquely identifies this Address.
	UUID string

	// Device is a reference to the device to which this Address is attached.
	Device *Device

	// IP is the IP address.
	IP net.IP

	// IPNet is the IP network of this address.
	IPNet *net.IPNet
}

// WithAddressUUID returns a function that sets the input UUID on an Address.
func WithAddressUUID(uuid string) func(*Address) {
	return func(a *Address) {
		a.UUID = uuid
	}
}

// OnDevice returns a function that sets the
// input Device reference on an Address.
func OnDevice(dev *Device) func(*Address) {
	return func(a *Address) {
		a.Device = dev
	}
}

// NewAddress returns a new Address based on the input arguments,
// or an error if an invalid address is supplied.
func NewAddress(addr string, options ...func(*Address)) (*Address, error) {
	var (
		ip    net.IP
		ipNet *net.IPNet
		err   error
	)

	if strings.Contains(addr, "/") {
		if ip, ipNet, err = net.ParseCIDR(addr); err != nil {
			return nil, err
		}
	} else {
		ip = net.ParseIP(addr)
	}

	a := &Address{
		IP:    ip,
		IPNet: ipNet,
	}

	for _, opt := range options {
		opt(a)
	}

	return a, nil
}
