package nexus

import (
	"net"
	"testing"
)

func TestHasIP(t *testing.T) {
	addr, err := NewAddress("192.168.0.128/24")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	dev := &Device{
		Addresses: []*Address{addr},
	}

	if !dev.HasIP(addr.IP) {
		t.Fatalf("expected Device with IP %v", addr.IP)
	}

	notIP := net.ParseIP("192.168.0.64")
	if dev.HasIP(notIP) {
		t.Fatalf("expected Device without IP %v", notIP)
	}
}
