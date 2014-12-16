package main

import (
	"fmt"
	"net"
	"os"
	"container/list"
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

	for { //forever loop handling requests
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
	parsed := parseCode(buf, bufflen)
	fmt.Println(parsed)
}

func parseCode(buf []byte, bufflen int) *list.List{

	ii := 0
	mystring := make([]byte, 512)
	var mainList *list.List = list.New()
	var curList *list.Element

	for i := 0; i < bufflen; i++ {
		switch string(buf[i]) {
		case "(":
			//at beginning of list, make a new list in the main list
			curList = mainList.PushBack(list.New())
		case ")":
			//at end of list, set the value in buffer to val of current list
			//curList.Value.PushBack(mystring[ii:bufflen])
			var curVal *list.List = curList.Value.(*list.List)
			curVal.PushBack(mystring[ii:bufflen])
			ii = 0;
		case "\n", "\r":
			//if new line character, do nothing
		default:
			mystring[ii] = buf[i]
			ii++
		}
	}

	return mainList
}
