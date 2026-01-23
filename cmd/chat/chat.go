package main

import (
	"bufio"
	db "chatRPC/db/rpc/clientStub"
	"chatRPC/dlog"
	myMessage "chatRPC/lib/message"
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

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func send(msg string) {
	myId := nodesetManager.GetId()
	me := nodesetManager.GetNode(myId)
	for _, node := range nodesetManager.GetNodeSet() {
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

	dlog.Printf("My address is %s\n", transport.GetAddress())

	//Bind chat to all the services endpoints
	db.Bind(os.Args[1])
	wg.Add(1)
	go func() {
		defer wg.Done()
		transport.Listen()
	}()
	nodeset.Bind(db.Get("nodeset"))

	//Register Chat Services
	nodemanager.Register()
	messenger.Register()

	fmt.Println("Welcome to chatRPC!")
	var username string
	for {
		fmt.Printf("Please enter your name: ")
		reader := bufio.NewReader(os.Stdin)
		name, err := reader.ReadString('\n')
		username = strings.TrimRight(name, "\r\n")
		if err != nil {
			panic(err)
		}

		if username != "" {
			break
		}
	}

	dlog.Printf("My username is: %s\n", username)
	nodesetManager.CreateCluster(username)

	//Render TUI
	app := tview.NewApplication()

	greetings := tview.NewTextView().
		SetText(fmt.Sprintf("Welcome to the chat room, %s!", username))
	greetings.SetBorder(true).
		SetTitle("chatRPC").
		SetTitleAlign(tview.AlignLeft)

	people := tview.NewTextView()
	messages := tview.NewTextView().SetDynamicColors(true)
	for _, node := range nodesetManager.GetNodeSet() {
		fmt.Fprintln(people, node.UserName)
	}

	nodesetManager.GetCluster().OnChange = func(diff *nodesetManager.DiffCluster) {
		app.QueueUpdateDraw(func() {
			for _, node := range diff.AddedNodes {
				fmt.Fprintln(people, node.UserName)
				fmt.Fprintf(messages, "[green::i]%s has entered the chat![-::-]\n", node.UserName)
			}

			if len(diff.RemovedNodes) > 0 {
				people.Clear()
				for _, node := range nodesetManager.GetCluster().NodeSet {
					fmt.Fprintln(people, node.UserName)
				}
				for _, node := range diff.RemovedNodes {
					fmt.Fprintf(messages, "[red::i]%s has left the chat![-::-]\n", node.UserName)
				}
			}

			diff.AddedNodes = diff.AddedNodes[:0]
			diff.RemovedNodes = diff.RemovedNodes[:0]
		})
	}

	myMessage.GetMessage().OnChange = func(msg *myMessage.Msg) {
		app.QueueUpdateDraw(func() {
			fmt.Fprintf(messages, "%s: %s\n", msg.From.UserName, msg.Message)
		})
	}

	people.SetBorder(true).
		SetTitle("People").
		SetTitleAlign(tview.AlignLeft)
	messages.SetBorder(true).
		SetTitle("Messages").
		SetTitleAlign(tview.AlignLeft)
	input := tview.NewInputField()

	input.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			me := nodesetManager.GetNode(nodesetManager.GetId())
			if input.GetText() == "" {
				return
			}
			fmt.Fprintf(messages, "%s: %s\n", me.UserName, input.GetText())
			send(input.GetText())
			input.SetText("")
		}
	})

	input.SetLabel("> ").
		SetFieldWidth(0).
		SetFieldBackgroundColor(tcell.ColorDefault)

	inputBox := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(input, 1, 0, true)

	inputBox.SetBorder(true).
		SetTitle("Type a message").
		SetTitleAlign(tview.AlignLeft)

	footer := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText("[yellow]Tab:[white]Type ↔ Scroll   [yellow]J/K:[white]Scroll down/up   [yellow]Ctrl+C:[white]Exit   [yellow]Enter:[white]Send message")

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(greetings, 3, 0, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(messages, 0, 3, false).
			AddItem(people, 25, 0, false), 0, 4, false).
		AddItem(inputBox, 7, 0, true).
		AddItem(footer, 1, 0, false)

	app.SetRoot(flex, true).SetFocus(input)

	focusOnInput := true
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			if focusOnInput {
				app.SetFocus(messages)
			} else {
				app.SetFocus(input)
			}
			focusOnInput = !focusOnInput
			return nil
		}

		return event
	})

	if err := app.Run(); err != nil {
		panic(err)
	}

	nodeset.Delete(nodesetManager.GetId())
	cancel()
	os.Exit(1)

}

var wg sync.WaitGroup
