package serverStub

import (
	"chatRPC/lib/transport"
	"chatRPC/nodeset"
	"encoding/json"
)

func Register() {
	transport.RegisterServerStub("Add", Add)
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
