package main

import (
	db "chatRPC/db/rpc/clientStub"
	"chatRPC/lib/nodesetManager"
	serverStub "chatRPC/lib/nodesetManager/rpc/serverStub"
	"chatRPC/lib/transport"
	nodeset "chatRPC/nodeset/rpc/clientStub"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {

	if len(os.Args) != 2 {
		panic(fmt.Errorf("usage %s <DBServerAddr>", os.Args[0]))
	}
	//add this node to the nodeset
	log.Printf("My address is %s\n", transport.GetAddress())

	//bind chat to all the services endpoints
	db.Bind(os.Args[1])
	wg.Add(1)
	go func() {
		defer wg.Done()
		transport.Listen()
	}()
	nodeset.Bind(db.Get("nodeset"))
	//message.Bind(db.Get("message"))

	//chatd stuff
	serverStub.Register()

	nodesetManager.CreateCluster()

	//while loop here
	// for {
	// 	fmt.Print("Type your message: ")
	// 	reader := bufio.NewReader(os.Stdin)
	// 	line, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if line == "quit\n" || line == "q\n" {
	// 		break
	// 	}

	// 	//send the message to the cluster
	// 	msg := message.Send(line)
	// 	fmt.Println(msg)
	// }

	wg.Wait()

	//call the clientStub
	// msg := message.Send(os.Args[2])
	// fmt.Println(msg)
}

var wg sync.WaitGroup
