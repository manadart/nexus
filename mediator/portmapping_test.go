package mediator

import (
	"github.com/manadart/nexus"
	"testing"
)

func TestNewPortMappingSimpleNoIPs(t *testing.T) {
	// Nexus providing service access on port 6380.
	nex := &nexus.Nexus{}
	nEP, err := nexus.NewEndpoint(6380)
	if err != nil {
		t.Fatal(err)
	}
	if err = nex.AddEndpoint(nEP); err != nil {
		t.Fatal(err)
	}

	// Service target listening to port 6379.
	target := &nexus.Nexus{}
	tEP, err := nexus.NewEndpoint(6379)
	if err != nil {
		t.Fatal(err)
	}
	if err = target.AddEndpoint(tEP); err != nil {
		t.Fatal(err)
	}

	cfg := PortMappingConfig{
		Name:       "Redis service port mapping",
		Nexus:      nex,
		Port:       6380,
		Targets:    []*nexus.Nexus{target},
		TargetPort: 6379,
	}
	pm, err := NewPortMapping(cfg)
	if err != nil {
		t.Fatal(err)
	}

	if pm == nil {
		t.Fatal("expected non-nil PortMapping")
	}
	if pm.Endpoint != nEP {
		t.Errorf("incorrect port")
	}
	if len(pm.Targets) != 1 {
		t.Error("expected 1 target Endpoint")
	}
	// TODO: Test other fields.
}
