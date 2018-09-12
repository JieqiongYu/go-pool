package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	listener, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		log.Fatal("Error starting tcp server:", err)
	}
	defer listener.Close()
	log.Println("Listening on " + connHost + ":" + connPort)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting:", err.Error())
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	fmt.Print("Message Received:", string(message))
	conn.Write([]byte(message + "\n"))
	conn.Close()
}
