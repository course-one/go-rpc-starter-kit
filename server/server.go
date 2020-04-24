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

	// register `mit` object with `rpc.DefaultServer`
	rpc.Register(mit)

	// register an HTTP handler for RPC communication on `http.DefaultServeMux` (default)
	// registers a handler on the `rpc.DefaultRPCPath` endpoint to respond to RPC messages
	// registers a handler on the `rpc.DefaultDebugPath` endpoint for debugging
	rpc.HandleHTTP()

	// sample test endpoint
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "RPC SERVER LIVE!")
	})

	// listen and serve default HTTP server
	http.ListenAndServe(":9000", nil)

}
