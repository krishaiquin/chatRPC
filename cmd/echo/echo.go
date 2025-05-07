package main

import (
	message "chatRPC/message/rpc/clientStub"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		panic(fmt.Errorf("usage %s <serverAddr> <message>", os.Args[0]))
	}
	//bind the server and the client
	message.Bind(os.Args[1])
	//call the clientStub
	msg := message.Echo(os.Args[2])
	fmt.Println(msg)
}
