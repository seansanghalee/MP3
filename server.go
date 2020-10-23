package main

import (
	"MP3_2/config"
	"encoding/gob"
	"fmt"
	"net"
)

func getID(c net.Conn) string {
	var ID string
	dec := gob.NewDecoder(c) // read from the channel
	dec.Decode(&ID)
	return ID
}

func main() {

	var (
		N, f int    // N: Total number of nodes, f: Upper bound for the number of faulty nodes
		port string // Port number the server listens to
	)

	N, f, port = config.ConfigureServer()

	nodes := make(map[string]net.Conn) // key: ID of the node, value: TCP connection with that node
	// var values []float64               // list to store the states of all the nodes

	// connects with all the processes
	for i := 1; i < N+1; i++ {
		l, err := net.Listen("tcp", ":"+port)
		if err != nil {
			fmt.Println(err)
			return
		}

		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		var temp string
		dec := gob.NewDecoder(c)
		dec.Decode(&temp)
		nodes[temp] = c
		l.Close()
	}

	fmt.Println(nodes)
	fmt.Println(f)

	for {

	}

	// start := time.Now()

	// for key, value := range nodes {
	// 	var message.Message
	// 	values := append(values, value)
	// }

	// // receive updates values
	// // for i := 0; i < N; i++ {

	// // }

	// // check if the states are within 0.01
	// //checkState(values)

	// // if done, end timer
	// elapsed := time.Since(start)
	// fmt.Println(elapsed)

}
