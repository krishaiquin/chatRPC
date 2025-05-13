package serverStub

import (
	"chatRPC/lib/transport"
	"chatRPC/message"
	"encoding/json"
)

func Register() {
	transport.RegisterServerStub("Send", Send)
}

func Send(args []byte) []byte {
	//unmarshal
	var msg string
	err := json.Unmarshal(args, &msg)
	if err != nil {
		panic(err)
	}
	//call the procedure
	res := message.Send(msg)
	//marshall
	data, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	return data
}
