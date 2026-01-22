package message

import (
	"chatRPC/nodeset/api"
	"sync"
)

type Msg struct {
	From     api.Node
	Message  string
	mx       sync.Mutex
	OnChange func(*Msg)
}

func Send(node api.Node, msg string) {
	message.mx.Lock()
	message.From = node
	message.Message = msg
	message.mx.Unlock()

	if message.OnChange != nil {
		message.OnChange(message)
	}

}

func GetMessage() *Msg {
	return message
}

func init() {
	message = &Msg{}
}

var message *Msg
