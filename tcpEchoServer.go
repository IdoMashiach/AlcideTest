package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func echo(conn net.Conn) {
	r := bufio.NewReader(conn)
	for {

		line, err := r.ReadBytes(byte('\n'))
		if err != nil {
			fmt.Println("ERROR", err)
			os.Exit(1) 
		}
		conn.Write(line)		
	}
}

func main() {

	//listen to the server
	l, err := net.Listen("tcp", "0.0.0.0:9001")
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

	//accept requests from client
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}
		go echo(conn)
	}
}











