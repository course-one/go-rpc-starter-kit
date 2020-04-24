package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/course-one/go-rpc-starter-kit/common"
)

func main() {

	// create a `*College` object
	mit := common.NewCollege()

	// create a custom RPC server
	server := rpc.NewServer()

	// register `mit` object with `rpc.DefaultServer`
	server.Register(mit)

	// create a TCP listener at address : 127.0.0.1:9002
	// https://golang.org/pkg/net/#Listener
	listener, _ := net.Listen("tcp", "127.0.0.1:9002")

	for {

		// get connection from the listener when client connects
		conn, _ := listener.Accept() // Accept blocks until next connection is received

		// serve connection in a separate goroutine using JSON codec
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
