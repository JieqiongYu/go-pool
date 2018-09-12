package main

import (
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
		log.Fatal("Error starting tcp server: ", err)
	}
	defer listener.Close()
	log.Println("Listening on " + connHost + ":" + connPort)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		log.Println(conn)
	}
}
