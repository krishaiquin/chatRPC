package serverStub

import (
	"chatRPC/lib/transport"
	"chatRPC/nodeset"
	"encoding/json"
)

func Register() {
	transport.RegisterServerStub("Add", Add)
	transport.RegisterServerStub("GetNodes", GetNodes)
}

func Add(addr []byte) []byte {
	//unmarshal
	var address string
	err := json.Unmarshal(addr, &address)
	if err != nil {
		panic(err)
	}
	//call the procedure
	nodeset.Add(address)

	return nil
}

func GetNodes(args []byte) []byte {

	cluster := nodeset.GetNodes()
	//marshal
	reply, err := json.Marshal(cluster)
	if err != nil {
		panic(err)
	}

	return reply
}
