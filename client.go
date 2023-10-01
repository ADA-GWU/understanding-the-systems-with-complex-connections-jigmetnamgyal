package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: client <server_address> <server_port>")
		os.Exit(1)
	}

	serverAddr := os.Args[1]
	serverPort := os.Args[2]

	for {
		fmt.Println("Enter a number: ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error Scanning: ", err)
			os.Exit(1)
		}

		conn, err := net.Dial("tcp", serverAddr+":"+serverPort)
		if err != nil {
			fmt.Println("Error connecting to the server:", err)
			continue
		}

		// Send input to the server
		_, err = conn.Write([]byte(input))

		if err != nil {
			fmt.Println("Error sending data to the server:", err)
			err := conn.Close()
			if err != nil {
				fmt.Println("Error closing", err)
				os.Exit(1)
			}
			continue
		}

		// Receive and print the result from the server
		result := make([]byte, 1024)
		_, err = conn.Read(result)
		if err != nil {
			fmt.Println("Error receiving data from the server:", err)
			err := conn.Close()
			if err != nil {
				fmt.Println("Error closing", err)
				os.Exit(1)
			}
			continue
		}

		fmt.Println(strings.TrimSpace(string(result)))

		ConnErr := conn.Close()
		if ConnErr != nil {
			fmt.Println("Error closing", ConnErr)
			os.Exit(1)
		}
	}
}
