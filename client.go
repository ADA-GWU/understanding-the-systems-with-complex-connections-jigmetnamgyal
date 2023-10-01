package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type ServerList struct {
	ServerAddress string
	ServerPort    string
}

func main() {

	fmt.Println("List the no of server: ")

	var noOfServer string
	_, err := fmt.Scanln(&noOfServer)

	if err != nil {
		fmt.Println("Error Scanning: ", err)
		os.Exit(1)
	}

	convertedInput, err := strconv.Atoi(noOfServer)

	if err != nil {
		fmt.Println("Error in converting string to int: ", err)
		os.Exit(1)
	}

	var serverLists []ServerList

	for i := 0; i < convertedInput; i++ {
		fmt.Println("List the server address: ")

		var ServerAddress string
		_, addressErr := fmt.Scanln(&ServerAddress)

		if addressErr != nil {
			fmt.Println("Error Scanning: ", addressErr)
			os.Exit(1)
		}

		fmt.Println("List the server port: ")

		var ServerPort string
		_, portErr := fmt.Scanln(&ServerPort)

		if portErr != nil {
			fmt.Println("Error Scanning: ", portErr)
			os.Exit(1)
		}

		serverLists = append(serverLists, ServerList{
			ServerAddress,
			ServerPort,
		})
	}

	fmt.Println(serverLists)

	for {
		fmt.Println("Enter a number: ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error Scanning: ", err)
			os.Exit(1)
		}

		randServer := returnRandomServer(serverLists)

		conn, err := net.Dial("tcp", randServer.ServerAddress+":"+randServer.ServerPort)
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

func returnRandomServer(serverLists []ServerList) ServerList {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := r.Intn(len(serverLists))

	randomServer := serverLists[randomIndex]

	return randomServer
}
