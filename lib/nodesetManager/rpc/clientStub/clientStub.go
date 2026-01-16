package clientStub

import (
	"chatRPC/lib/transport"
	"chatRPC/nodeset/api"
	"encoding/json"
)

func AddMember(destination string, node api.Node) {

	data, err := json.Marshal(node)
	if err != nil {
		panic(err)
	}

	transport.Call(destination, "AddMember", data)
}

func RemoveMember(destination string, nodeId uint32) {

	data, err := json.Marshal(nodeId)
	if err != nil {
		panic(err)
	}

	transport.Call(destination, "RemoveMember", data)
}
