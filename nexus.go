package nexus

// Nexus represents a network entity to which connections can be made.
type Nexus struct {
	// UUID uniquely identifies this Nexus.
	UUID string

	// Devs are the link-layer devices attached to this entity.
	Devs []*Device

	// Endpoints are the service endpoints being offered by this entity.
	Endpoints []*Endpoint
}
