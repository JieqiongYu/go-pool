package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

/*Args is struct holds information about argumetns passed from the client(RPC) to the server.*/
type Args struct{}

/*TimeServer is a number to register with the rpc.Register*/
type TimeServer int64

/*GiveServerTime is the function that will be called by the client, and the current server time is returned.*/
// GiveServerTime takes the Args object as the first argument and a reply pointer object.
// It sets the reply pointer object but does not return anything except an error
// The Args struct here has no fields because this server is not expecting the client to send any arguments
func (t *TimeServer) GiveServerTime(arg *Args, reply *int64) error {
	// Fill reply pointer to send the data back
	*reply = time.Now().Unix()
	return nil
}

func main() {
	// Create a new RPC server
	timeserver := new(TimeServer)
	// Register RPC server
	rpc.Register(timeserver)
	// HandleHTTP registers an HTTP handler for RPC messages to DefaultServer.
	rpc.HandleHTTP()
	// Listen for requests on port 1234
	// Start a TCP server that listens on port 1234.
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	// http.Serve is used to serve it as a running program.
	http.Serve(l, nil)
}
