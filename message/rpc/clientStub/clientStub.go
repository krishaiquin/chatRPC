package clientStub

import (
	"chatRPC/lib/transport"
	"encoding/json"
)

// bind communication line to serverAddress
func Bind(addr string) {
	server = addr
}

func Send(message string) string {
	//marshal
	data, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	//send it to server and get a response
	msg := transport.Call(server, "Send", data)
	var response string
	//unmarshal
	err = json.Unmarshal(msg, &response)
	if err != nil {
		panic(err)
	}
	//return
	return response

}

var server string
