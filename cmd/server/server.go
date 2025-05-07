package main

import (
	"chatRPC/lib/transport"
	"fmt"
)

func main() {

	fmt.Printf("Endpoint: %s\n", transport.GetAddress())
	transport.Listen()
}
