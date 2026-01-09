package clientStub

import (
	"chatRPC/lib/transport"
	"encoding/json"
	"log"
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
	log.Printf("Add ClientStub received response: %s", response)
	//unmarshal response
	var nodeId uint32
	err = json.Unmarshal(response, &nodeId)
	if err != nil {

		panic(err)
	}

	return nodeId
}

func Delete(nodeId uint32) {
	args, err := json.Marshal(nodeId)
	if err != nil {
		panic(err)
	}

	response := transport.Call(server, "Delete", args)
	log.Printf("Delete ClientStub received response: %s", response)
}

var server string
