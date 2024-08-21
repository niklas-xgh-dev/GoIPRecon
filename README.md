# go-ip-recon

A simple IP reconnaissance tool written in Go. This project aims to gather basic information about an IP address using various network tools.

## Features

- OS detection via ping TTL analysis
- Hostname lookup using nslookup
- Basic port scanning for the 20 most important ports
- Windows compatibility
- Traceroute analysis

## Requirements

- Go 1.16 or higher
- Windows OS (for full functionality)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/niklas-xgh-dev/go-ip-recon.git
   ```

2. Navigate to the project directory:
   ```
   cd go-ip-recon
   ```

3. Build the project:
   ```
   go build ./cmd/main.go
   ```

## Usage

Run the compiled binary:

```
./main
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
```

## Traceroute Analysis

The traceroute analysis feature maps the network path to the target IP address. It provides information about:

- Number of hops to reach the destination
- IP addresses of intermediate routers
- Response times for each hop

This feature helps in understanding the network topology and identifying potential bottlenecks or points of interest in the route to the target IP.

## Contributing

This is a hobby project. If you have any cool ideas or improvements, let me know!

## Disclaimer

This tool is for educational purposes only. Ensure you have permission before scanning any IP addresses or networks you do not own.

## License

[MIT License](LICENSE.md)