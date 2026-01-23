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

	data, err := json.Marshal(args)
	if err != nil {
		panic(err)
	}

	transport.Call(to, "Send", data)
}
