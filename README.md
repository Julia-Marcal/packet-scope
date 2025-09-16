packet-scope
GoShark is a lightweight network packet monitoring and analysis tool written in Go. It demonstrates how to separate domain logic from infrastructure using Domain-Driven Design (DDD) principles, while providing basic packet sniffing and analysis features.

```
packet-scope/
├── cmd/
│   └── monitor/
│       └── main.go                  # Application entry point
├── internal/
│   ├── domain/
│   │   └── model/
│   │       └── packet.go            # Entities and value objects
│   ├── application/
│   │   └── analyzer.go              # Packet analysis logic
│   ├── infrastructure/
│   │   └── repository/
│   │       └── memory/
│   │           └── packet_repository.go  # In-memory repository implementation
│   ├── config/
│   │   └── config.go                # Configuration loading
│   ├── interface/
│   │   ├── logger/
│   │   │   └── logger.go            # Logging configuration
│   │   └── handler/
│   │       └── packet_handler.go    # HTTP or CLI handlers
├── pkg/
│   └── utils/
│       └── utils.go                 # Shared utility functions
├── config/
│   └── config.yaml                  # Configuration file
├── go.mod
├── go.sum
└── README.md
```

### 🧠 Folder Overview

- **/cmd/monitor**: Application entry point. Initializes and configures all required components.
- **/internal/domain**: Domain model definitions, including entities (e.g., `Packet`) and value objects.
- **/internal/application**: Business logic, such as packet analysis and alert generation.
- **/internal/infrastructure/repository**: Data access implementations, e.g., in-memory storage for packets.
- **/internal/config**: Loads and manages application configuration (network parameters, analysis rules, etc).
- **/internal/interface/logger**: Logging system configuration.
- **/internal/interface/handler**: Exposes application functionality via HTTP, CLI, or other protocols.
- **/pkg/utils**: Shared utility functions.
- **/config**: Configuration files (YAML/JSON) for parameters like packet size limits or monitored protocols.

---

### ⚠️ Troubleshooting: `panic: couldn't load wpcap.dll`

To capture network packets on Windows, you need Npcap (recommended) or WinPcap (obsolete).  
**Download and install Npcap:** [https://nmap.org/npcap/](https://nmap.org/npcap/)
