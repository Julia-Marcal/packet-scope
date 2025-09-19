
# Packet Scope

A lightweight network packet monitoring and analysis tool written in Go. It demonstrates clean architecture principles with proper separation of concerns, while providing real-time packet sniffing and analysis features.

## 📁 Project Structure

```
packet-scope/
├── cmd/
│   └── monitor/
│       └── main.go                          # Application entry point
├── config/
│   └── config.yaml                          # Configuration file
├── internal/
│   ├── application/
│   │   └── analyzer.go                      # Application service layer
│   ├── config/
│   │   └── config.go                        # Configuration management
│   ├── domain/
│   │   └── model/
│   │       └── packet.go                    # Domain entities and value objects
│   └── infrastructure/
│       └── capture/
│           └── gopacket_capture.go          # Packet capture implementation
├── pkg/
│   ├── logger/
│   │   └── logger.go                        # Colored logging system
│   └── utils/
│       └── utils.go                         # Network utility functions
├── vendor/                                  # Vendored dependencies
├── go.mod                                   # Go module definition
├── go.sum                                   # Dependency checksums
├── LICENSE                                  # Project license
└── README.md                                # Project documentation
```

## 🧠 Folder Overview

- **`/cmd/monitor`**: Application entry point that initializes and starts the packet analysis service.
- **`/internal/application`**: Application service layer containing business logic and use cases for packet analysis.
- **`/internal/domain/model`**: Domain entities and value objects including packet data structures and capture configurations.
- **`/internal/infrastructure/capture`**: Infrastructure layer with packet capture implementation using gopacket library.
- **`/internal/config`**: Configuration management for loading and handling application settings.
- **`/pkg/logger`**: Colored logging system with multiple log levels (Info, Success, Warning, Error, Process, System).
- **`/pkg/utils`**: Shared utility functions for network operations like local IP detection.
- **`/config`**: Configuration files (YAML) for application parameters and settings.
- **`/vendor`**: Vendored dependencies managed by Go modules.

## 🚀 Features

- **Real-time packet capture** on multiple network interfaces
- **Clean architecture** following Domain-Driven Design principles
- **Colored logging** with multiple severity levels
- **Local traffic filtering** to monitor only outgoing packets
- **Protocol analysis** for TCP, UDP, and other network protocols
- **Cross-platform support** with proper Windows network capture support

## 🛠️ Prerequisites

- Go 1.21.5 or higher
- Network capture library (Npcap for Windows, libpcap for Unix-like systems)
- Administrator/root privileges for packet capture

## 📦 Installation

1. **Install Npcap** (Windows) or ensure libpcap is available (Linux/macOS)
   - Windows: Download from [https://nmap.org/npcap/](https://nmap.org/npcap/)

2. **Clone the repository**
   ```bash
   git clone https://github.com/Julia-Marcal/packet-scope.git
   cd packet-scope
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Build the application**
   ```bash
   go build -o packet-scope ./cmd/monitor
   ```

## 🏃‍♂️ Usage

Run the packet monitor with administrator privileges:

```bash
# Windows (run as Administrator)
.\packet-scope.exe

# Linux/macOS (run with sudo)
sudo ./packet-scope
```

The application will:
- Detect all available network interfaces
- Start capturing packets on each interface
- Filter and display only outgoing traffic from your machine
- Show packet information including IP addresses, protocols, and ports

## 📊 Example Output

```
[15:04:05] [INFO] Starting packet capture on 3 interfaces with local IP: 192.168.1.100
[eth0] 15:04:05 | Len=54 | IP 192.168.1.100 -> 8.8.8.8 | TCP 45678 -> 53
[wlan0] 15:04:06 | Len=42 | IP 192.168.1.100 -> 192.168.1.1 | UDP 68 -> 67
```

## ⚠️ Troubleshooting

### Windows: `panic: couldn't load wpcap.dll`

To capture network packets on Windows, you need Npcap (recommended) or WinPcap (obsolete).

**Solution:**
1. Download and install Npcap: [https://nmap.org/npcap/](https://nmap.org/npcap/)
2. During installation, make sure to check "Install Npcap in WinPcap API-compatible mode"
3. Restart your computer after installation
4. Run the application as Administrator

### Linux/macOS: Permission Denied

If you get permission errors when trying to capture packets:

**Solution:**
```bash
# Run with sudo
sudo ./packet-scope

# Or give the binary packet capture capabilities (Linux only)
sudo setcap cap_net_raw,cap_net_admin=eip ./packet-scope
```

### No Interfaces Found

If the application reports "no network interfaces found":

**Solutions:**
- Ensure you're running with administrator/root privileges
- Check that your network capture library is properly installed
- Verify that network interfaces are active and available

## 🏗️ Architecture

This project follows Clean Architecture principles:

- **Domain Layer** (`internal/domain`): Contains business entities and rules
- **Application Layer** (`internal/application`): Contains use cases and application services  
- **Infrastructure Layer** (`internal/infrastructure`): Contains external implementations
- **Interface Layer** (`pkg/logger`, `cmd`): Contains external interfaces and entry points

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes following conventional commits
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
