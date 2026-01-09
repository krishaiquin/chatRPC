package serverStub

import (
	"chatRPC/lib/transport"
	"chatRPC/nodeset"
	"encoding/json"
)

func Register() {
	transport.RegisterServerStub("Add", Add)
	transport.RegisterServerStub("Delete", Delete)
}

func Add(args []byte) []byte {
	var addr string
	err := json.Unmarshal(args, &addr)
	if err != nil {
		panic(err)
	}

	//call the routine
	nodeId := nodeset.Add(addr)

	//marshal the result
	data, err := json.Marshal(nodeId)
	if err != nil {
		panic(err)
	}

	return data
}

func Delete(args []byte) []byte {
	var nodeId uint32
	err := json.Unmarshal(args, &nodeId)
	if err != nil {
		panic(err)
	}

	nodeset.Delete(nodeId)

	return nil

}
