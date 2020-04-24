package main

import (
	"io"
	"net/http"
	"net/rpc"

	"github.com/course-one/go-rpc-starter-kit/common"
)

func main() {

	// create a `*College` object
	mit := common.NewCollege()

	// create a custom RPC server
	server := rpc.NewServer()

	// register `mit` object with `rpc.DefaultServer`
	server.Register(mit)

	// register an HTTP handler for RPC communication on `http.DefaultServeMux` (default)
	// '/rpc' => for client-server communication
	// '/rpc-debug' => for debugging
	server.HandleHTTP("/rpc", "/rpc-debug")

	// sample test endpoint
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "RPC SERVER LIVE!")
	})

	// listen and serve default HTTP server
	http.ListenAndServe(":9001", nil)

}
