package main



import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleServerConn(clientConn, 5*time.Second)
	}
}

func handleServerConn(client net.Conn, timeout time.Duration) {
	err := client.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		fmt.Println("Error setting read deadline:", err)
	}

	var size uint32
	err = binary.Read(client, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}

	err = client.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		fmt.Println("Error setting read deadline:", err)
	}

	bytMsg := make([]byte, size)
	_, err = client.Read(bytMsg)
	if err != nil {
		panic(err)
	}
	strMsg := string(bytMsg)
	fmt.Printf("%s\n", strMsg)

	reply := "Message received"

	err = client.SetWriteDeadline(time.Now().Add(timeout))
	if err != nil {
		fmt.Println("Error setting write deadline:", err)
	}

	err = binary.Write(client, binary.LittleEndian, uint32(len(reply)))
	if err != nil {
		panic(err)
	}

	err = client.SetWriteDeadline(time.Now().Add(timeout))
	if err != nil {
		fmt.Println("Error setting write deadline:", err)
	}

	_, err = client.Write([]byte(reply))
	if err != nil {
		panic(err)
	}
}
