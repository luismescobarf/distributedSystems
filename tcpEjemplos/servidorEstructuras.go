// server.go

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

const MIN = 1
const MAX = 100

func random() int {
	return rand.Intn(MAX-MIN) + MIN
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {

		//_, err := conn.Read(tmp)
		//conn.Read(tmp)

		tmp := make([]byte, 500)

		_, err := c.Read(tmp)
		if err != nil {
			fmt.Println(err)
			return
		}

		// convert bytes into Buffer (which implements io.Reader/io.Writer)
		tmpbuff := bytes.NewBuffer(tmp)

		tmpstruct := new(Message)

		// creates a decoder object
		gobobj := gob.NewDecoder(tmpbuff)

		// decodes buffer and unmarshals it into a Message struct
		gobobj.Decode(tmpstruct)

		// lets print out!
		fmt.Println(tmpstruct) // reflects.TypeOf(tmpstruct) == Message{}

		//c.Write([]byte(string(result)))

		/*
			/////////////////
			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}

			temp := strings.TrimSpace(string(netData))
			if temp == "STOP" {
				break
			}

			result := strconv.Itoa(random()) + "\n"
			fmt.Println("Escribiendo")
			c.Write([]byte(string(result)))
		*/
	}
	fmt.Println("Cerrando conexión!")
	//c.Close()
	c.Close()
}

// Create your custom data struct
type Message struct {
	ID   string
	Data string
}

/*
func main() {

	// for purpose of verbosity, I will be removing error handling from this
	// sample code

	server, _ := net.Listen("tcp", ":12345")
	conn, _ := server.Accept()

	// create a temp buffer
	tmp := make([]byte, 500)

	// loop through the connection to read incoming connections. If you're doing by
	// directional, you might want to make this into a seperate go routine
	for {

		//_, err := conn.Read(tmp)
		conn.Read(tmp)

		// convert bytes into Buffer (which implements io.Reader/io.Writer)
		tmpbuff := bytes.NewBuffer(tmp)

		tmpstruct := new(Message)

		// creates a decoder object
		gobobj := gob.NewDecoder(tmpbuff)

		// decodes buffer and unmarshals it into a Message struct
		gobobj.Decode(tmpstruct)

		// lets print out!
		fmt.Println(tmpstruct) // reflects.TypeOf(tmpstruct) == Message{}

	}

}
*/

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	//PORT = ":12345"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
