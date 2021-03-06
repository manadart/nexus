package nexus

import (
	"net"
	"testing"
)

func TestNexusHasIP(t *testing.T) {
	addr, err := NewAddress("192.168.0.128/24")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	dev := &Device{Addresses: []*Address{addr}}
	nex := &Nexus{Devices: []*Device{dev}}

	if !nex.HasIP(addr.IP) {
		t.Fatalf("expected Device with IP %v", addr.IP)
	}

	notIP := net.ParseIP("192.168.0.64")
	if nex.HasIP(notIP) {
		t.Fatalf("expected Device without IP %v", notIP)
	}
}
