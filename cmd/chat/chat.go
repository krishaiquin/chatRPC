package main

import (
	db "chatRPC/db/rpc/clientStub"
	"chatRPC/lib/transport"
	message "chatRPC/message/rpc/clientStub"
	nodeset "chatRPC/nodeset/rpc/clientStub"
)

func main() {

	// if len(os.Args) != 2 {
	// 	panic(fmt.Errorf("usage %s <serverAddr>", os.Args[0]))
	// }

	//bind chat to all the services endpoints
	nodeset.Bind(db.Get("nodeset"))
	message.Bind(db.Get("message"))

	//add this node to the nodeset
	nodeset.Add(transport.GetAddress())

	//while loop here

	// //call the clientStub
	// msg := message.Echo(os.Args[2])
	// fmt.Println(msg)
}
