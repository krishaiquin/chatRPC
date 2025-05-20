package main

import (
	db "chatRPC/db/rpc/clientStub"
	"chatRPC/lib/transport"
	"chatRPC/message/rpc/serverStub"
	nodeset "chatRPC/nodeset/rpc/clientStub"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Errorf("usage %s <DB_Server_Addr>", os.Args[0]))
	}
	//bind to db server. set the destination address for db requests from message server
	db.Bind(os.Args[1])
	nodeset.Bind(db.Get("nodeset"))
	//register message service functions
	serverStub.Register()
	//create an endpoint for message service
	fmt.Printf("Listening on: %s\n", transport.GetAddress())
	//save this service endpoint to db
	db.Put("message", transport.GetAddress())

	transport.Listen()
}
