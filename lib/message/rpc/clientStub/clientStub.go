package clientStub

import (
	"chatRPC/lib/message/api"
	"chatRPC/lib/transport"
	"encoding/json"
)

func Send(to string, from uint32, message string) string {
	args := api.SendArgs{
		From:    from,
		Message: message,
	}
	//marshal
	data, err := json.Marshal(args)
	if err != nil {
		panic(err)
	}
	//send it to server and get a response
	msg := transport.Call(to, "Send", data)
	var response string
	//unmarshal
	err = json.Unmarshal(msg, &response)
	if err != nil {
		panic(err)
	}
	//return
	return response

}
