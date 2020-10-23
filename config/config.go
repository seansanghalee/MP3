package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Configure() (int, int, []string, []string, []string) {
	file := ReadConfig()
	N, f, IDs, IPs, ports := Extract(file)
	display(N, f, IDs, IPs, ports)
	return N, f, IDs, IPs, ports
}

func ConfigureServer() (int, int, string) {
	file := ReadConfig()
	N, f, _, _, ports := Extract(file)
	serverPort := ports[0]
	return N, f, serverPort
}

func display(N, f int, IDs, IPs, ports []string) {
	fmt.Println("---Configuration File Read---")
	fmt.Printf("N: %v\n", N)
	fmt.Printf("f: %v\n", f)
	fmt.Printf("IDs: %v\n", IDs)
	fmt.Printf("IPs: %v\n", IPs)
	fmt.Printf("Ports: %v\n", ports)
	fmt.Println("-----------------------------")
}

func Extract(str []string) (int, int, []string, []string, []string) {
	values := strings.Split(str[0], " ")
	N, err := strconv.Atoi(values[0])
	if err != nil {
		fmt.Println(err)
	}
	f, err := strconv.Atoi(values[1])
	if err != nil {
		fmt.Println(err)
	}

	var IDs, IPs, ports []string

	for i := 1; i < len(str); i++ {
		temp := strings.Split(str[i], " ")
		IDs = append(IDs, temp[0])
		IPs = append(IPs, temp[1])
		ports = append(ports, temp[2])
	}

	return N, f, IDs, IPs, ports
}

func GetID(port string, ports []string, IDs []string) string {
	var ID string

	for i := 0; i < len(ports); i++ {
		if ports[i] == port {
			ID = IDs[i]
		}
	}

	return ID
}

func GetPort() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the port number: ")
	port, _ := reader.ReadString('\n')
	port = strings.TrimSpace(port)

	return port
}

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
