package serverStub

import (
	"chatRPC/lib/nodesetManager"
	"chatRPC/lib/transport"
	"chatRPC/nodeset/api"
	"encoding/json"
)

func Register() {
	transport.RegisterServerStub("AddMember", AddMember)
	transport.RegisterServerStub("RemoveMember", RemoveMember)

}

func AddMember(data []byte) []byte {
	var node api.Node

	err := json.Unmarshal(data, &node)
	if err != nil {
		panic(err)
	}

	nodesetManager.AddMember(node)

	return nil
}

func RemoveMember(data []byte) []byte {
	var nodeId uint32

	err := json.Unmarshal(data, &nodeId)
	if err != nil {
		panic(err)
	}

	nodesetManager.RemoveMember(nodeId)

	return nil
}
