package serverStub

import (
	"chatRPC/lib/transport"
	"chatRPC/nodeset"
	"encoding/json"
)

func Register() {
	transport.RegisterServerStub("Add", Add)
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
