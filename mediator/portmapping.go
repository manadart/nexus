package mediator

import (
	"errors"
	"fmt"
	"net"

	"github.com/manadart/nexus"
)

// PortMapping represents a port mapping arrangement between a Nexus instance
// and one or more target Nexus instances. It can be used to represent:
// - Port mapping
// - Port forwarding
// - Load balancing
type PortMapping struct {
	// UUID uniquely identifies this PortMapping.
	UUID string

	// Name describes this mapping. E.g. "API load balancer".
	Name string

	// Endpoint is the service access Endpoint for this PortMapping.
	Endpoint *nexus.Endpoint

	// Targets are the service Endpoints to which traffic
	// is being forwarded by this PortMapping.
	Targets []*nexus.Endpoint

	// Protocol is the protocol (TCP or UDP) being forwarded
	// over this port mapping. It can be unspecified.
	Protocol string
}

// PortMappingConfig is the config required to create a PortMapping.
type PortMappingConfig struct {
	Name     string
	Protocol string

	Nexus *nexus.Nexus
	Port  int
	IP    net.IP

	Targets    []*nexus.Nexus
	TargetPort int
	TargetIPs  []net.IP
}

// Validate ensures that the PortMappingConfig describes a PortMapping
// that is actually possible between the input Nexus and Targets.
// If mappings are declared for services listening on specific IPs,
// those IPs must be assigned to the corresponding Nexus instances.
// TODO: Add tests for this.
func (cfg PortMappingConfig) Validate() error {
	if cfg.Nexus == nil {
		return errors.New("port mapping requires a nexus")
	}
	if cfg.Port == 0 {
		return errors.New("port mapping requires a port")
	}

	if len(cfg.Targets) == 0 {
		return errors.New("port mapping requires one or more targets")
	}
	if cfg.TargetPort == 0 {
		return errors.New("port mapping requires a target port")
	}

	// Ensure that a populated IP is represented by one of the Nexus' addresses.
	if cfg.IP != nil && !cfg.IP.IsUnspecified() {
		if !cfg.Nexus.HasIP(cfg.IP) {
			return fmt.Errorf("nexus does not have address %s", cfg.IP)
		}
	}

	tCount := len(cfg.Targets)
	ipCount := len(cfg.TargetIPs)

	// If populated, the input TargetIPs must match the number of Targets.
	if ipCount != 0 {
		if ipCount != tCount {
			return fmt.Errorf("number of targets and target IPs must match; got %d and %d", tCount, ipCount)
		}
		for i, ip := range cfg.TargetIPs {
			if ip != nil && !ip.IsUnspecified() {
				if !cfg.Targets[i].HasIP(ip) {
					return fmt.Errorf("target at index %d does not have address %s", i, cfg.IP)
				}
			}
		}
	}

	// TODO: Check for endpoint conflicts with IP etc.

	return nil
}

// NewPortMapping returns new PortMapping based on the input configuration,
// or an error if the configuration is invalid.
func NewPortMapping(cfg PortMappingConfig) (*PortMapping, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	ep := cfg.Nexus.EndpointForPort(cfg.Port)
	// TODO: Add if not extant?
	if ep == nil {
		return nil, fmt.Errorf("nexus %s has no endpoint for port %d", cfg.Nexus.UUID, cfg.Port)
	}

	targets := make([]*nexus.Endpoint, len(cfg.Targets))
	for i, target := range cfg.Targets {
		if targets[i] = target.EndpointForPort(cfg.TargetPort); targets[i] == nil {
			return nil, fmt.Errorf("nexus %s has no endpoint for port %d", target.UUID, cfg.TargetPort)
		}
	}

	pm := &PortMapping{
		UUID:     "",
		Name:     cfg.Name,
		Endpoint: ep,
		Targets:  targets,
		Protocol: cfg.Protocol,
	}

	return pm, nil
}
