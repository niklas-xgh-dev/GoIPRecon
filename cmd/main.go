package main

import (
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var relevantPorts = []int{21, 22, 23, 25, 53, 80, 110, 135, 139, 143, 443, 445, 1433, 3306, 3389, 5900, 8080, 8443, 5060, 6379}

func main() {
	var ipAddress string
	fmt.Print("Enter an IP address: ")
	fmt.Scanln(&ipAddress)

	fmt.Println("\n--- IP Reconnaissance Results ---")

	// OS Detection
	os := detectOS(ipAddress)
	fmt.Printf("Operating System: %s\n", os)

	// Hostname Lookup
	hostname := lookupHostname(ipAddress)
	fmt.Printf("Hostname: %s\n", hostname)

	// Port Scanning
	openPorts := scanPorts(ipAddress)
	fmt.Println("Open ports:", openPorts)

	// Traceroute Analysis
	performTraceroute(ipAddress)
}

func detectOS(ip string) string {
	cmd := exec.Command("ping", "-n", "1", "-w", "1000", ip)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	outputStr := string(output)
	if strings.Contains(outputStr, "TTL=128") {
		return "Windows"
	} else if strings.Contains(outputStr, "TTL=64") {
		return "Linux"
	} else if strings.Contains(outputStr, "TTL=255") {
		return "Unix/FreeBSD"
	}
	return "Unknown"
}

func lookupHostname(ip string) string {
	hostnames, err := net.LookupAddr(ip)
	if err != nil || len(hostnames) == 0 {
		return "Unknown"
	}
	return hostnames[0]
}

func scanPorts(ip string) []int {
	openPorts := []int{}
	for _, port := range relevantPorts {
		address := fmt.Sprintf("%s:%d", ip, port)
		conn, err := net.DialTimeout("tcp", address, 200*time.Millisecond)
		if err == nil {
			openPorts = append(openPorts, port)
			conn.Close()
		}
	}
	return openPorts
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
