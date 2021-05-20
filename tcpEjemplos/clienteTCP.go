package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:6666")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs))

	/*

		strEcho := "Halo"
		servAddr := "localhost:6666"
		tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
		if err != nil {
			println("ResolveTCPAddr failed:", err.Error())
			os.Exit(1)
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			println("Dial failed:", err.Error())
			os.Exit(1)
		}


		_, err = conn.Write([]byte(strEcho))
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		println("write to server = ", strEcho)

		reply := make([]byte, 1024)

		fmt.Println("Pasé por acá!")

		_, err = conn.Read(reply)
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		fmt.Println("reply from server=", string(reply))

		conn.Close()
	*/
}
