package main

import (
	db "chatRPC/db/rpc/clientStub"
	"chatRPC/lib/transport"
	message "chatRPC/message/rpc/clientStub"
	nodeset "chatRPC/nodeset/rpc/clientStub"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		panic(fmt.Errorf("usage %s <DBServerAddr> <message>", os.Args[0]))
	}

	//bind chat to all the services endpoints
	db.Bind(os.Args[1])
	nodeset.Bind(db.Get("nodeset"))
	message.Bind(db.Get("message"))

	//add this node to the nodeset
	nodeset.Add(transport.GetAddress())

	//while loop here

	//call the clientStub
	msg := message.Send(os.Args[2])
	fmt.Println(msg)
}
