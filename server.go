package main

import (
	"MP3_2/config"
	"MP3_2/helper"
	"MP3_2/message"
	"MP3_2/receiver"
	"MP3_2/sender"
	"fmt"
	"net"
	"time"
)

func main() {
	N, _, port := config.ConfigureServer() // N: Total number of nodes, f: Upper bound for the number of faulty nodes, port: port number the server listens to
	nodes := make(map[string]net.Conn)     // key: ID of the node, value: TCP connection with that node

	// connects with all the processes
	for i := 1; i < N+1; i++ {
		receiver.ServerListen(port, nodes)
	}

	helper.DisplayMap(nodes)

	start := time.Now() // start timer

	var r int

	// while states are not within 0.01 keep receiving progress from nodes
	for {
		values := []float64{} // a list to store the states of all the nodes

		for key, value := range nodes {
			var toReceive message.Message
			receiver.UnicastReceive(value, &toReceive)
			if toReceive.State == 0 {
				delete(nodes, key)
				fmt.Println("Process", key, "crashed")
			} else {
				values = append(values, toReceive.State)
				r = toReceive.Round
			}
		}

		helper.PrintServerRoundInfo(r, values)

		if helper.CheckState(values) {
			break
		}
	}

	elapsed := time.Since(start) // end timer when done

	sender.SendExit(nodes, r)
	fmt.Println("It took", r, "rounds and", elapsed, "to reach consensus.")
	return
}
