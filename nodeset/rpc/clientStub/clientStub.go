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
	transport.Call(server, "Add", data)

}

func GetNodes() []string {

	reply := transport.Call(server, "GetNodes", nil)
	var res []string
	//unmarshal
	err := json.Unmarshal(reply, &res)
	if err != nil {
		panic(err)
	}

	return res

}

var server string
