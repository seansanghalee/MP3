package sender

import (
	"MP3_2/message"
	"encoding/gob"
	"fmt"
	"net"
)

// Dial does what
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

// UnicastSend sends
func UnicastSend(destination net.Conn, message message.Message) {
	// rand.Seed(time.Now().UnixNano())
	// duration := rand.Float64()
	// time.Sleep(time.Duration(duration) * time.Millisecond)

	//set delay
	groupTest := new(sync.WaitGroup)
	go delay(minDelay, maxDelay, groupTest)

	//Wait group is used to block the execution of code in the main thread until all goroutines are complete and waitgroup counter is decremented to 0
	groupTest.Add(1)
	groupTest.Wait()

	enc := gob.NewEncoder(destination)
	enc.Encode(message)
}


func delay(min int, max int, wg *sync.WaitGroup) {
	num := rand.Intn(max-min) + min
	time.Sleep(time.Duration(num) * time.Millisecond)

	//decrement value of waitgroup and relay the flow of execution back to main
	wg.Done()
}
