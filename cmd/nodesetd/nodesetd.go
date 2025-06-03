package main

import (
	db "chatRPC/db/rpc/clientStub"
	"chatRPC/lib/transport"
	"chatRPC/nodeset/rpc/serverStub"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Errorf("usage %s <DB_Server_Addr>", os.Args[0]))
	}
	//bind to db server. set the destination address for db requests from nodeset server
	db.Bind(os.Args[1])
	//Register nodeset function
	serverStub.Register()
	fmt.Printf("Listening on: %s\n", transport.GetAddress())
	db.Put("nodeset", transport.GetAddress())
	transport.Listen()
}
