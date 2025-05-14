package main

import (
	serverStub "chatRPC/db/rpc/serverStub"
	"chatRPC/lib/transport"
	"fmt"
)

func main() {
	serverStub.Register()
	fmt.Printf("Listening on: %s\n", transport.GetAddress())
	transport.Listen()
}
