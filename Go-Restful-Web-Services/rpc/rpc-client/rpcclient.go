package main

import (
	"log"
	"net/rpc"
)

/*Args is the struct for the Arguments passed to the server.*/
type Args struct {
}

func main() {
	var reply int64
	args := Args{}
	// Do a DialHTTP to connect to the RPC server, which is running on the localhost on port 1234.
	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Call the Remote function with the `Name:Function` format with args and reply with the pointer object.
	// Get the data collected into the reply object.
	// The Call function is sequential in nature
	err = client.Call("TimeServer.GiveServerTime", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Printf("%d", reply)
}
