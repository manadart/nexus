package nexus

import (
	"testing"
)

func TestNewEndpoint(t *testing.T) {
	port := 80

	ep, err := NewEndpoint(port)
	if err != nil {
		t.Fatal(err)
	}
	if ep.Port != port {
		t.Errorf("expected Endpoint with port %d, got %d", port, ep.Port)
	}

	// Options.
	id := "123-456-789"
	ep, err = NewEndpoint(port, WithEndpointUUID(id), OnNexus(&Nexus{}))
	if err != nil {
		t.Fatal(err)
	}
	if ep.UUID != id {
		t.Errorf("expected UUID %s, got %s", id, ep.UUID)
	}
	if ep.Nexus == nil {
		t.Error("expected non-nil Nexus")
	}
}
