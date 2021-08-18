# Nexus

Nexus is a Go library for modelling networks.

## Basic Entities

The network entities are modelled using simple types.

### Nexus

Some presence on the network - a machine, VM, container etc.

### Device

A link-layer device (real or virtual) connected to a Nexus.

### Address

An IP address assigned to a Device.

### Endpoint

A service/socket offered by a Nexus listening to a port.

### Model

The entity relationships are represented in the following diagram.

```
┌──────────────────────────────────────────────────────────────────┐
│                                                                  │
│ Nexus (Machine/VM/Container)                                     │
│                                                                  │
│                                                                  │
│                                                                  │
│  ┌────────────────────────────┐  ┌────────────────────────────┐  │
│  │ Endpoint (Port 80)         │  │ Endpoint (Port 6379)       │  │
│  │                            │  │                            │  │
│  │ IP 192.168.0.10            │  │ IP None/0.0.0.0/::         │  │
│  │               │            │  │ (Listening on all)         │  │
│  │               │            │  │                            │  │
│  │               │            │  │                            │  │
│  │               │            │  │                            │  │
│  │               │            │  │                            │  │
│  │               │            │  │                            │  │
│  └───────────────┼────────────┘  └────────────────────────────┘  │
│                  │                                               │
│  ┌───────────────┼────────────┐  ┌────────────────────────────┐  │
│  │ Device (eth0) │            │  │ Device (eth1)              │  │
│  │               │            │  │                            │  │
│  │               │            │  │                            │  │
│  │ ┌─────────────▼────┐       │  │ ┌───────────────────┐      │  │
│  │ │ Address          │       │  │ │ Address           │      │  │
│  │ │ (192.168.0.10)   │       │  │ │ (10.0.0.128)      │      │  │
│  │ │                  │       │  │ │                   │      │  │
│  │ │                  │       │  │ │                   │      │  │
│  │ └──────────────────┘       │  │ └───────────────────┘      │  │
│  │                            │  │                            │  │
│  └────────────────────────────┘  └────────────────────────────┘  │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

## Relationships

Relationships between entities in the model are represented using the [mediator
pattern](https://en.wikipedia.org/wiki/Mediator_pattern).

The following relationships are available.

### PortMapping

Used to represent:
- Port forwarding
- Port mapping
- Load balancing

It connects an Endpoint to one or more target Endpoints.

```
                                              ┌─────────────────────────┐
                                              │ Nexus (Service pod)     │
                                              │                         │
                                              │ ┌─────────────────────┐ │
                                       ┌──────┼─► Endpoint (Port 80)  │ │
                                       │      │ └─────────────────────┘ │
                                       │      │                         │
                                       │      └─────────────────────────┘
┌───────────────────────┐  ┌───────────┼────┐
│                       │  │           │    │ ┌─────────────────────────┐
│ Nexus (Load balancer) │  │ Port      │    │ │ Nexus (Service pod)     │
│                       │  │ Mapping   │    │ │                         │
│ ┌───────────────────┐ │  │           │    │ │ ┌─────────────────────┐ │
│ │ Endpoint (Port 80)├─┼──┼───────────┼────┼─┼─► Endpoint (Port 80)  │ │
│ │                   │ │  │           │    │ │ └─────────────────────┘ │
│ └───────────────────┘ │  │           │    │ │                         │
│                       │  │           │    │ └─────────────────────────┘
└───────────────────────┘  └───────────┼────┘
                                       │      ┌─────────────────────────┐
                                       │      │ Nexus (Service pod)     │
                                       │      │                         │
                                       │      │ ┌─────────────────────┐ │
                                       └──────┼─► Endpoint (Port 80)  │ │
                                              │ └─────────────────────┘ │
                                              │                         │
                                              └─────────────────────────┘
```

### HostContainer (TBC)

Used to represent a host-container association via a bridge network.

### Others...
