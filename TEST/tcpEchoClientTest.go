package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
) 

//Explanation:
//when i came ro write this test i had few limitations.
//   i didnt succeeded to test the process of getting user input and check if 
//   the print on the console is the same.
//   so, i did something little different, but its preety much imitate the process. 

func client(data *string) {

	conn, err := net.Dial("tcp", "0.0.0.0:9001")
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
	
	
	response := bufio.NewReader(conn)
	for {
		userLine := []byte(*data)
		conn.Write(userLine)
		
		serverLine, err := response.ReadBytes(byte('\n'))
		if err == nil {
			*data = string(serverLine)
			_ = *data //  just for not to yelling at me. not so elegant.
		} else { 
			fmt.Println("ERROR", err) 
		}
	}
}



























