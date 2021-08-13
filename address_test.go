package nexus

import (
	"testing"
)

func TestNewAddress(t *testing.T) {
	// Plain IP.
	ip := "192.168.0.5"
	addr, err := NewAddress(ip)
	if err != nil {
		t.Error(err)
	}
	if addr.IP == nil || addr.IP.String() != ip {
		t.Errorf("expected IP %s, got %s", ip, addr.IP )
	}

	// IP with mask.
	ipSub := "192.168.0.5/24"
	addr, err = NewAddress(ipSub)
	if err != nil {
		t.Error(err)
	}
	if addr.IP == nil || addr.IP.String() != ip {
		t.Errorf("expected IP %s, got %s", ip, addr.IP )
	}
	if addr.IPNet == nil || addr.IPNet.String() != "192.168.0.0/24" {
		t.Errorf("expected IPNet %s, got %s", "192.168.0.0/24", addr.IPNet )
	}

	// Options.
	id := "123-456-789"
	addr, err = NewAddress("", WithUUID(id), OnDevice(&Device{}))
	if err != nil {
		t.Error(err)
	}
	if addr.UUID != id {
		t.Errorf("expected UUID %s, got %s", id, addr.UUID)
	}
	if addr.Device == nil {
		t.Error("expected non-nil Device")
	}
}
