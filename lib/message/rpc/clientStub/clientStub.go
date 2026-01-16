package clientStub

import (
	messageAPI "chatRPC/lib/message/api"
	"chatRPC/lib/transport"
	nodesetAPI "chatRPC/nodeset/api"
	"encoding/json"
)

func Send(to string, from nodesetAPI.Node, message string) {
	args := messageAPI.SendArgs{
		From:    from,
		Message: message,
	}
	//marshal
	data, err := json.Marshal(args)
	if err != nil {
		panic(err)
	}
	//send it to server and get a response
	transport.Call(to, "Send", data)
}
