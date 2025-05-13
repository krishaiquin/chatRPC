package main

import (
	"chatRPC/lib/transport"
	"fmt"
)

func main() {
	fmt.Printf("Listening on: %s\n", transport.GetAddress())
	transport.Listen()
}
