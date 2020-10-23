package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Configure reads the config file and returns the contents in seperate files
func Configure() (int, int, []string, []string, []string, []string) {
	file := ReadConfig()
	N, f, IDs, IPs, ports, faultyOrNot := Extract(file)
	display(N, f, IDs, IPs, ports)
	return N, f, IDs, IPs, ports, faultyOrNot
}

// ConfigureServer reads the config file and returns information for server node to start
func ConfigureServer() (int, int, string) {
	file := ReadConfig()
	N, f, _, _, ports, _ := Extract(file)
	serverPort := ports[0]
	return N, f, serverPort
}

// Display displays the config file in a readable format
func display(N, f int, IDs, IPs, ports []string) {
	fmt.Println("---Configuration File Read---")
	fmt.Printf("N: %v\n", N)
	fmt.Printf("f: %v\n", f)
	fmt.Printf("IDs: %v\n", IDs)
	fmt.Printf("IPs: %v\n", IPs)
	fmt.Printf("Ports: %v\n", ports)
	fmt.Println("-----------------------------")
}

// Extract takes the []string config file and returns the contents in seperate variables
func Extract(str []string) (int, int, []string, []string, []string, []string) {
	values := strings.Split(str[0], " ")
	N, err := strconv.Atoi(values[0])
	if err != nil {
		fmt.Println(err)
	}
	f, err := strconv.Atoi(values[1])
	if err != nil {
		fmt.Println(err)
	}

	var IDs, IPs, ports, faultyOrNot []string

	for i := 1; i < len(str); i++ {
		temp := strings.Split(str[i], " ")
		IDs = append(IDs, temp[0])
		IPs = append(IPs, temp[1])
		ports = append(ports, temp[2])
		faultyOrNot = append(faultyOrNot, temp[3])
	}

	return N, f, IDs, IPs, ports, faultyOrNot
}

// GetIDFromPort gets the ID of a process from the port it is connected to
func GetIDFromPort(port string, ports []string, IDs []string) string {
	var ID string

	for i := 0; i < len(ports); i++ {
		if ports[i] == port {
			ID = IDs[i]
		}
	}

	return ID
}

// GetFaultyFromPort gets whether the node is failty or not from the port it is connected to
func GetFaultyFromPort(port string, ports []string, faultyOrNot []string) bool {
	faulty := false

	for i := 0; i < len(ports); i++ {
		if ports[i] == port {
			if faultyOrNot[i] == "f" {
				faulty = true
			}
		}
	}

	return faulty
}

// ReadConfig reads config.txt and returns it as a []string
func ReadConfig() []string {
	file, err := os.Open("config.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	var str []string

	for scanner.Scan() {
		str = append(str, scanner.Text())
	}

	file.Close()
	return str
}
