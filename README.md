# Nexus

Nexus is a Go library for modelling networks.

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
