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
	"strings"
	"sync"
	"syscall"

	"github.com/rivo/tview"
)

func send(msg string) {
	myId := nodesetManager.GetId()
	me := nodesetManager.GetNode(myId)
	for _, node := range nodesetManager.GetCluster() {
		if node.NodeId == myId {
			continue
		}
		message.Send(node.Addr, me, msg)
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

	fmt.Println("Welcome to the chat room!")
	fmt.Printf("Please enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	username = strings.TrimRight(username, "\r\n")
	if err != nil {
		panic(err)
	}

	dlog.Printf("My username is: %s\n", username)
	nodesetManager.CreateCluster(username)

	//TUI stuff
	app := tview.NewApplication()

	input := tview.NewInputField().
		SetLabel("> ").
		SetFieldWidth(0)

	// Then wrap it in a bordered box if you want a border:
	inputBox := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(input, 1, 0, true)

	inputBox.SetBorder(true).
		SetTitle("Type a message").
		SetTitleAlign(tview.AlignLeft)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Welcome to the chat room!").SetTitleAlign(tview.AlignLeft), 4, 0, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Messages").SetTitleAlign(tview.AlignLeft), 0, 3, false).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("People").SetTitleAlign(tview.AlignLeft), 25, 0, false), 0, 4, false).
			AddItem(inputBox, 7, 0, true), 0, 4, false)
	if err := app.SetRoot(flex, true).SetFocus(input).Run(); err != nil {
		panic(err)
	}

	for {
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
		dlog.Printf("Sending line: %s\n", line)
		send(line)

	}
}

var wg sync.WaitGroup
