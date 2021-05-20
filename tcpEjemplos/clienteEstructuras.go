package main

import (
	"bytes"
	"encoding/gob"
	"net"
	"os"
)

type Message struct {
	ID   string
	Data string
}

func main() {
	// lets create the message we want to send accross
	msg := Message{ID: "Yo", Data: "Hello"}
	bin_buf := new(bytes.Buffer)

	// error handling still truncated
	servAddr := "localhost:12345"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	// create a encoder object
	gobobj := gob.NewEncoder(bin_buf)

	// encode buffer and marshal it into a gob object
	gobobj.Encode(msg)

	conn.Write(bin_buf.Bytes())

	conn.Close()
}
