package nexus

// Device represents a physical or virtual network device that is addressable
// at the link-layer via a hardware address, such as MAC or InfiniBand GUID.
type Device struct {
	// UUID uniquely identifies this device.
	UUID string

	// NexusUUID identifies the Nexus to which this device is attached.
	NexusUUID string

	// HWAddr is the link-layer address of the device.
	HWAddr string

	// NetNSID identifies the network namespace for this device.
	// See https://en.wikipedia.org/wiki/Linux_namespaces#Network_(net).
	NetNSID string

	// Addresses are the IP addresses attached to this device.
	Addresses []Address
}
