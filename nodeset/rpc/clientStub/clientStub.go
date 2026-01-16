package clientStub

import (
	"chatRPC/dlog"
	"chatRPC/lib/transport"
	"chatRPC/nodeset/api"
	"encoding/json"
)

func Bind(addr string) {
	server = addr
}

func Add(addr string, username string) (uint32, []api.Node) {
	args := api.AddArgs{
		Addr:     addr,
		Username: username,
	}
	data, err := json.Marshal(args)
	if err != nil {
		panic(err)
	}

	//send it to nodeset server
	response := transport.Call(server, "Add", data)

	dlog.Printf("Add Nodeset ClientStub received response: %s", response)
	//unmarshal response
	var retVal api.AddRet
	err = json.Unmarshal(response, &retVal)
	if err != nil {
		panic(err)
	}

	return retVal.NodeId, retVal.NodeSet
}

func Delete(nodeId uint32) {
	args, err := json.Marshal(nodeId)
	if err != nil {
		panic(err)
	}

	response := transport.Call(server, "Delete", args)
	dlog.Printf("Delete ClientStub received response: %s", response)
}

var server string
