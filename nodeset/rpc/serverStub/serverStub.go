package serverStub

import (
	"chatRPC/nodeset"
	"encoding/json"
)

func Add(addr []byte) []byte {
	//unmarshal
	var address string
	err := json.Unmarshal(addr, &address)
	if err != nil {
		panic(err)
	}
	//call the procedure
	res := nodeset.Add(address)
	//marshall
	data, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	return data
}
