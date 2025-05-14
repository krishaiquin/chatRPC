package clientStub

import (
	"chatRPC/db/api"
	"chatRPC/lib/transport"
	"encoding/json"
)

func Bind(addr string) {
	server = addr
}
func Put(service string, endpoint string) string {
	args := api.PutArgs{
		Service:  service,
		Endpoint: endpoint,
	}
	data, err := json.Marshal(args)
	if err != nil {
		panic(err)
	}
	response := transport.Call(server, "Put", data)
	var result string
	err = json.Unmarshal(response, &result)
	if err != nil {
		panic(err)
	}

	return result
}

var server string
