package main

import (
	"MP3_2/config"
	"MP3_2/helper"
	"MP3_2/message"
	"MP3_2/receiver"
	"MP3_2/sender"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	var (
		N, f                         int      // N: Total number of nodes, f: Upper bound for the number of faulty nodes
		IDs, IPs, ports, faultyOrNot []string // IDs: Process IDs of nodes, IPs/ports: IPs/ports for processes
		port, ID                     string   // ID: ID of this process, port: port number of this process
		faulty                       = false  // faulty: true if faulty node, false otherwise
	)

	N, f, IDs, IPs, ports, faultyOrNot = config.Configure()

	//port = config.GetPort()
	port = os.Args[1]
	ID = config.GetIDFromPort(port, ports, IDs)
	faulty = config.GetFaultyFromPort(port, ports, faultyOrNot)

	nodes := make(map[string]net.Conn) // create a map to store connections to other nodes {key: process id, value: TCP connection}

	IDinInt, _ := strconv.Atoi(ID)

	// establishes connection

	go receiver.Listen(port, IDinInt, IDs, nodes) // listens to all the other nodes that will be dialing this process

	for i := 0; i < IDinInt; i++ {
		sender.Dial(i, ID, IDs, IPs, ports, nodes)
	} // dials all the other nodes that has started already and been listening for this process

	for len(nodes) < N {
	} // blocks until all nodes are connected

	helper.DisplayMap(nodes)
	if faulty {
		fmt.Println("THIS IS A FAULTY NODE")
	}

	y, r := helper.Initialize()

	go receiver.ListenForExit(nodes)

	for {
		states := []float64{}

		for key, value := range nodes {
			if key != "0" { // don't send to server or
				toSend := message.Message{y, r}
				sender.UnicastSend(value, toSend)
			}
		} // multicast_send

		for key, value := range nodes {
			if key != "0" { // don't receive from server
				toReceive := message.Message{}
				receiver.UnicastReceive(value, &toReceive)
				//fmt.Println("Received y:", toReceive.State, ", r:", toReceive.Round, "from", key)
				if toReceive.Round == r {
					states = append(states, toReceive.State)
				}
			}
		} // multicast_receive

		states = append(states, y)
		helper.PrintRoundInfo(y, r, states)

		// update y and r
		y = helper.Average(states[0 : N-f])
		fmt.Println("***Updated y: ", y)

		// send server the updated value
		m := message.Message{y, r}
		sender.UnicastSend(nodes["0"], m)
		r++

		if faulty {
			helper.NodeCrash()
		}
	}
}
