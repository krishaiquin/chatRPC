package main

import (
	"chatRPC/lib/transport"
	nodeset "chatRPC/nodeset/rpc/clientStub"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		panic(fmt.Errorf("usage %s <serverAddr>", os.Args[0]))
	}
	//bind the server and the client
	nodeset.Bind(os.Args[1])
	//add this node to the nodeset
	nodeset.Add(transport.GetAddress())

	// //call the clientStub
	// msg := message.Echo(os.Args[2])
	// fmt.Println(msg)
}
