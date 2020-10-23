package sender

import (
	"MP3_2/message"
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
)

// Dial calls the provided network address to create a TCP connection
func Dial(i int, ID string, IDs, IPs []string, ports []string, nodes map[string]net.Conn) {
	c, err := net.Dial("tcp", IPs[i]+":"+ports[i])
	if err != nil {
		fmt.Println(err)
		return
	}

	enc := gob.NewEncoder(c)
	enc.Encode(ID)
	nodes[IDs[i]] = c
}

// UnicastSend sends a message to other process via TCP channel
func UnicastSend(destination net.Conn, message message.Message) {
	groupTest := new(sync.WaitGroup) // block the execution of code in the main thread until all goroutines are complete and waitgroup counter is decremented to 0
	go delay(100, 200, groupTest)
	groupTest.Add(1)
	groupTest.Wait()

	rand.Seed(time.Now().UnixNano())
	duration := rand.Float64()
	time.Sleep(time.Duration(duration) * time.Second) // add delay to simulate real-life unicast_send

	enc := gob.NewEncoder(destination)
	enc.Encode(message)
}

// Delay decrements value of waitgroup and relay the flow of execution back to main
func delay(min int, max int, wg *sync.WaitGroup) {
	num := rand.Intn(max-min) + min
	time.Sleep(time.Duration(num) * time.Millisecond)

	wg.Done()
}

// SendExit sends exit signals to connected nodes
func SendExit(nodes map[string]net.Conn, r int) {
	for _, value := range nodes {
		m := message.Message{0, r}
		UnicastSend(value, m)
	}
}
