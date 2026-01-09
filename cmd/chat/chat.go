package main

import (
	"bufio"
	db "chatRPC/db/rpc/clientStub"
	"chatRPC/dlog"
	message "chatRPC/lib/message/rpc/clientStub"
	messenger "chatRPC/lib/message/rpc/serverStub"
	"chatRPC/lib/nodesetManager"
	nodemanager "chatRPC/lib/nodesetManager/rpc/serverStub"
	"chatRPC/lib/transport"
	nodeset "chatRPC/nodeset/rpc/clientStub"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func send(msg string) {
	my_id := nodesetManager.GetId()
	for _, node := range nodesetManager.GetCluster() {
		if node.NodeId == my_id {
			continue
		}
		message.Send(node.Addr, my_id, msg)
	}
}

func main() {

	if len(os.Args) != 2 {
		panic(fmt.Errorf("usage %s <DBServerAddr>", os.Args[0]))
	}
	ctx, cancel := context.WithCancel(context.Background())

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	defer func() {
		signal.Stop(signalCh)
		cancel()
	}()

	go func() {
		select {
		case <-signalCh:
			nodeset.Delete(nodesetManager.GetId())
			cancel()
			fmt.Printf("Cancellation Recieved! Exiting...\n")
			os.Exit(1)
		case <-ctx.Done():
		}
	}()
	//add this node to the nodeset
	dlog.Printf("My address is %s\n", transport.GetAddress())

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
	nodemanager.Register()
	messenger.Register()

	nodesetManager.CreateCluster()

	fmt.Println("Welcome to chatRPC!")

	for {
		fmt.Print("Type your message: ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if line == "quit\n" || line == "q\n" {
			nodeset.Delete(nodesetManager.GetId())
			cancel()
			fmt.Printf("Bye!\n")
			os.Exit(1)
		}
		send(line)

	}

}

var wg sync.WaitGroup
