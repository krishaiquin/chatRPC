package clientStub

import (
	"chatRPC/lib/transport"
	"encoding/json"
)

func Bind(addr string) {
	server = addr
}

func Add(addr string) uint32 {
	//marshal addr
	data, err := json.Marshal(addr)
	if err != nil {
		panic(err)
	}

	//send it to nodeset server
	response := transport.Call(server, "Add", data)

	//unmarshal response
	var nodeId uint32
	err = json.Unmarshal(response, &nodeId)
	if err != nil {
		panic(err)
	}

	return nodeId
}

var server string
