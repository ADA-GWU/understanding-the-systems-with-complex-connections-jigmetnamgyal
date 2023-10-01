package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	// Check if the correct number of command-line arguments is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: server <port>")
		os.Exit(1)
	}

	// Get the port number from the command-line argument
	port := os.Args[1]

	// Listen on the specified port for incoming connections
	listener, err := net.Listen("tcp", "localhost:"+port)

	if err != nil {
		fmt.Println("Error listening to port: ", err)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("Error closing the listener with error: ", err)
			return
		}
	}(listener)

	fmt.Println("Server Listening on port: " + port)

	for {
		// Accept incoming connections
		connection, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err)
			continue
		}

		// Handle each client connection in a separate goroutine
		go handleClient(port, connection)
	}
}

func handleClient(port string, conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing the listener with error: ", err)
			os.Exit(1)
		}
	}(conn)

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}

		data := string(buffer[:n])
		inputNumber, err := strconv.Atoi(data)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			_, err := conn.Write([]byte("Invalid input. Please enter a number.\n"))

			if err != nil {
				fmt.Println("Error writing:", err)
				return
			}
			continue
		}

		result := doubleNumber(inputNumber)
		_, writeError := conn.Write([]byte(fmt.Sprintf("Result from server port(%s) = %d\n", port, result)))

		if writeError != nil {
			fmt.Println("Error writing:", writeError)
			return
		}
	}
}

func doubleNumber(number int) int {
	return number * 2
}
