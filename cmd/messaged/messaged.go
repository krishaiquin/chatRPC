package main

import (
	"chatRPC/lib/transport"
	message "chatRPC/message/rpc/serverStub"
	"fmt"
)

func main() {
	message.Register()
	fmt.Printf("Listening on: %s\n", transport.GetAddress())
	transport.Listen()
}
