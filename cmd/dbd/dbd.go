package main

import (
	serverStub "chatRPC/db/rpc/serverStub"
	"chatRPC/lib/transport"
	"fmt"
	"sync"
)

func main() {
	serverStub.Register()
	fmt.Printf("Listening on: %s\n", transport.GetAddress())
	wg.Add(1)
	go func() {
		defer wg.Done()
		transport.Listen()
	}()

	//wait until goroutine is done
	wg.Wait()
}

var wg sync.WaitGroup
