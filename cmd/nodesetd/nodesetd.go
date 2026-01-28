package main

import (
	"chatRPC/lib/transport"
	"chatRPC/nodeset/rpc/serverStub"
	"fmt"
	"sync"
)

/**
*	Runs Nodeset server
 */
func main() {
	serverStub.Register()
	fmt.Printf("Listening on: %s\n", transport.GetAddress())
	wg.Add(1)
	go func() {
		defer wg.Done()
		transport.Listen()
	}()

	wg.Wait()
}

var wg sync.WaitGroup
