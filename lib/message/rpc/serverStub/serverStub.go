package serverStub

import (
	"chatRPC/dlog"
	"chatRPC/lib/message"
	"chatRPC/lib/message/api"
	"chatRPC/lib/transport"
	"encoding/json"
)

func Register() {
	transport.RegisterServerStub("Send", Send)
}

func Send(args []byte) []byte {
	//unmarshal
	var msg api.SendArgs
	err := json.Unmarshal(args, &msg)
	if err != nil {
		panic(err)
	}
	dlog.Printf("Received the message!\n")
	//call the procedure
	message.Send(msg.From, msg.Message)

	return nil
}
