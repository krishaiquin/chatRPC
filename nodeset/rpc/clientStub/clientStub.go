package clientStub

import (
	"chatRPC/lib/transport"
	"encoding/json"
)

func Bind(addr string) {
	server = addr
}

func Add(addr string) {
	//marshal
	data, err := json.Marshal(addr)
	if err != nil {
		panic(err)
	}
	//send it to server
	transport.Call(server, "add", data)

}

var server string
