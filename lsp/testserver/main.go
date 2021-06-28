package main

import (
	"context"
	"go.starlark.net/lsp/jsonrpc2"
	"log"
)

var addr = "localhost:10420"

type debugPreempter struct {}

func (d debugPreempter) Preempt(ctx context.Context, req *jsonrpc2.Request) (interface{}, error) {
	log.Printf("received message: %s\n", req)
	return nil, jsonrpc2.ErrNotHandled
}

func main() {
	ctx := context.Background()

	listener, err := jsonrpc2.NetListener(ctx, "tcp", addr, jsonrpc2.NetListenOptions{})
	if err != nil {
		log.Fatalf("error creating listener: %s", err)
	}


	server, err := jsonrpc2.Serve(ctx, listener, jsonrpc2.ConnectionOptions{
		Preempter: debugPreempter{},
	})
	if err != nil {
		log.Fatalf("error creating server: %s", err)
	}

	log.Printf("listening on %s\n", addr)
	if err := server.Wait(); err != nil {
		log.Fatalf("server exited with error: %s", err)
	}
}
