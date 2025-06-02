package main

import (
	"bufio"
	db "chatRPC/db/rpc/clientStub"
	"chatRPC/lib/nodesetManager"
	message "chatRPC/message/rpc/clientStub"
	nodeset "chatRPC/nodeset/rpc/clientStub"
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		panic(fmt.Errorf("usage %s <DBServerAddr> <message>", os.Args[0]))
	}

	//bind chat to all the services endpoints
	db.Bind(os.Args[1])
	nodeset.Bind(db.Get("nodeset"))
	message.Bind(db.Get("message"))

	//add this node to the nodeset
	nodesetManager.CreateCluster()

	//while loop here
	for {
		fmt.Print("Type your message: ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if line == "quit\n" || line == "q\n" {
			break
		}

		//send the message to the cluster
		msg := message.Send(line)
		fmt.Println(msg)
	}

	//call the clientStub
	// msg := message.Send(os.Args[2])
	// fmt.Println(msg)
}
