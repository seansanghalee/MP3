package receiver

import (
	"MP3_2/message"
	"encoding/gob"
	"fmt"
	"net"
)

// Listen does what
func Listen(port string, ID int, IDs []string, nodes map[string]net.Conn) {
	for {
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
}

// UnicastReceive receives
func UnicastReceive(source net.Conn, message *message.Message) {
	dec := gob.NewDecoder(source)
	dec.Decode(message)
}
