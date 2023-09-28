package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: client <server_address> <server_port>")
		os.Exit(1)
	}
}
