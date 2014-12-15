package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//Listen for incoming connections
	l, err := net.Listen("tcp", "localhost:3333")
	//Handle errors
	if err != nil {
		fmt.Println("Error Listening", err.Error())
		os.Exit(1)
	}
	//defer the close
	defer l.Close()
	fmt.Println("Listening on: ", 3333)

	for { //forever loop hdling requests
		//accept requests
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error on accept", err.Error())
			os.Exit(1)
		}
		//handle Request in new goroutine
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	//Make buffer as slice
	defer conn.Close()
	buf := make([]byte, 1024)
	//read into buffer
	bufflen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error on read: ", err.Error())
		os.Exit(1)
	}
	fmt.Printf(string(buf))
	fmt.Println(bufflen)
	conn.Write(buf)
}
