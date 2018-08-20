package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
) 

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:9001")
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

	userInput := bufio.NewReader(os.Stdin)
	response := bufio.NewReader(conn)
	for {
		userLine, err := userInput.ReadBytes(byte('\n'))
		
		if err == nil {
			conn.Write(userLine)
		} else { 
			fmt.Println("ERROR", err) 
		}

		serverLine, err := response.ReadBytes(byte('\n'))
		if err == nil {
			fmt.Print(string(serverLine))
		} else { 
			fmt.Println("ERROR", err) 
		}
	}
}




















