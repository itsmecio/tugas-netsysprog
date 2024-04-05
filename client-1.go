package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func menu() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("1.Send msg to server\n2. Keluar\n>>")
		scanner.Scan()
		s := scanner.Text()
		if s == "1" {
			WriteMsg()
		} else if s == "2" {
			fmt.Println("Exit")
			break
		}

	}
}

func WriteMsg() {
	scanner := bufio.NewScanner(os.Stdin)
	var msg string
	for {
		fmt.Print("enter message: ")
		scanner.Scan()
		Message = scanner.Text()
		if len(Message) < 1 {
			fmt.Println("Message contains minimal 1 character")
		} else {
			break
		}
	}
	SendMessage(Message)
}

func SendMessage(Message string) {
	serverConn, err := net.DialTimeout("tcp", "127.0.0.1:1234", 3*time.Second)
	if err != nil {
		panic(err)
	}
	defer serverConn.Close()

	err = serverConn.SetWriteDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		fmt.Println("Error setting write deadline:", err)
	}

	err = binary.Write(serverConn, binary.LittleEndian, uint32(len(Message)))
	if err != nil {
		panic(err)
	}

	err = serverConn.SetWriteDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		fmt.Println("Error setting write deadline:", err)
	}

	_, err = serverConn.Write([]byte(Message))
	if err != nil {
		panic(err)
	}

	err = serverConn.SetReadDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		fmt.Println("Error setting read deadline:", err)
	}

	var size uint32
	err = binary.Read(serverConn, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}

	err = serverConn.SetReadDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		fmt.Println("Error setting read deadline:", err)
	}

	bytReply := make([]byte, size)
	_, err = serverConn.Read(bytReply)
	if err != nil {
		panic(err)
	}
	fmt.Printf("replied: %s\n", string(bytReply))
}

func main() {
	menu()
}
