package serverStub

import (
	"chatRPC/lib/nodesetManager"
	"chatRPC/lib/transport"
	"chatRPC/nodeset/api"
	"encoding/json"
)

func Register() {
	transport.RegisterServerStub("Update", Update)
}

func Update(data []byte) []byte {

	//unmarshal args
	var args []api.Node
	err := json.Unmarshal(data, &args)
	if err != nil {
		panic(err)
	}

	//call the routine
	nodesetManager.Update(args)

	return nil
}
