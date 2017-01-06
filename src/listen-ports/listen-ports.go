package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(os.Args[0], " <start port> <end port>")
		os.Exit(1)
	}
	start := os.Args[1]
	end := os.Args[2]

	start_port, err := strconv.Atoi(start)
	if err != nil {
		panic(err)
	}

	end_port, err := strconv.Atoi(end)
	if err != nil {
		panic(err)
	}

	if end_port < start_port {
		fmt.Println("Error! End port is smaller than Start port.")
		os.Exit(1)
	}

	if end_port <= 0 || start_port <= 0 {
		fmt.Println("Port must be positive.")
		os.Exit(1)
	}

	for i := start_port; i <= end_port; i++ {
		listen(i)
	}

	for {
		// wait a Ctrl+C
	}
}

func listen(i int) {
	port := strconv.Itoa(i)
	fmt.Println("Listen port:", port)
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	go func() {
		defer ln.Close()
		for {
			conn, err := ln.Accept()
			if err == nil {
				fmt.Println(conn.LocalAddr())
			} else {
				fmt.Println(err)
				os.Exit(3)
			}
		}
	}()
}
