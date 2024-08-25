# go-ip-recon

A simple IP reconnaissance tool written in Go. This project aims to gather basic information about an IP address using various network tools.

## Features

- OS detection via ping TTL analysis
- Hostname lookup using nslookup
- Basic port scanning for 20 cybersecurity-relevant ports
- Cross-platform support (Windows and macOS)
- Traceroute analysis

## Requirements

- Go 1.16 or higher

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/niklas-xgh-dev/go-ip-recon.git
   ```

2. Navigate to the project directory:
   ```
   cd go-ip-recon
   ```

## Building

### For macOS:
```
./builds/build-mac.sh
```

### For Windows:
```
./builds/build-windows.sh
```

## Usage

### On macOS:
Run the compiled binary:
```
./builds/go-ip-recon
```

### On Windows:
Run the compiled binary:
```
.\builds\go-ip-recon.exe
```

Follow the prompts to enter an IP address for reconnaissance.

## Project Structure

```
go-ip-recon/
├── cmd/
│   └── main.go
├── internal/
│   ├── recon/
│   │   └── recon.go
│   └── osdetect/
│       └── osdetect.go
├── builds/
│   ├── build-mac.sh
│   ├── build-windows.sh
├── .gitignore
├── go.mod
├── README.md
└── LICENSE.md
```

## Features in Detail

### OS Detection
Estimates the target's operating system based on TTL values from ping responses.

### Hostname Lookup
Performs a reverse DNS lookup to find the hostname associated with the IP address.

### Port Scanning
Scans 20 cybersecurity-relevant ports, including common services like HTTP, SSH, FTP, etc.

### Traceroute Analysis
Maps the network path to the target IP address, providing information about:
- Number of hops to reach the destination
- IP addresses of intermediate routers
- Response times for each hop

## Contributing

This is a hobby project. If you have any cool ideas or improvements, feel free to fork the repository and submit a pull request!

## Disclaimer

This tool is for educational purposes only. Ensure you have permission before scanning any IP addresses or networks you do not own.

## License

[MIT License](LICENSE.md)
