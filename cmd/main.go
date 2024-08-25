package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/niklas-xgh-dev/go-ip-recon/internal/osdetect"
	"github.com/niklas-xgh-dev/go-ip-recon/internal/recon"
)

func main() {
	var ipAddress string
	fmt.Print("Enter an IP address: ")
	fmt.Scanln(&ipAddress)

	fmt.Println("\n--- IP Reconnaissance Results ---")

	// OS Detection
	osInfo, err := osdetect.DetectOS(ipAddress)
	if err != nil {
		log.Printf("Error detecting OS: %v", err)
	}
	fmt.Printf("Operating System: %s %s\n", osInfo.Name, osInfo.Version)
	if osInfo.IsVM {
		fmt.Println("This appears to be a virtual machine.")
	}

	// Reconnaissance
	result, err := recon.ScanIP(ipAddress)
	if err != nil {
		log.Printf("Error during reconnaissance: %v", err)
	}

	fmt.Printf("Hostname: %s\n", result.Hostname)
	fmt.Printf("Open ports: %v\n", result.OpenPorts)

	// Traceroute Analysis
	performTraceroute(ipAddress)
}

func performTraceroute(ip string) {
	fmt.Println("\nTraceroute Analysis:")
	cmd := exec.Command("tracert", "-d", "-w", "500", ip)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error performing traceroute:", err)
		return
	}

	lines := strings.Split(string(output), "\n")
	hopRegex := regexp.MustCompile(`\s*(\d+)\s+(\d+)\s+ms\s+(\d+)\s+ms\s+(\d+)\s+ms\s+(\S+)`)

	for _, line := range lines {
		matches := hopRegex.FindStringSubmatch(line)
		if len(matches) == 6 {
			hopNumber := matches[1]
			ipAddress := matches[5]
			fmt.Printf("Hop %s: %s\n", hopNumber, ipAddress)
		}
	}
}
