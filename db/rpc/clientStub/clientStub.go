package clientStub

import (
	"chatRPC/db/api"
	"chatRPC/lib/transport"
	"encoding/json"
)

func Bind(addr string) {
	server = addr
}
func Put(service string, endpoint string) {
	args := api.PutArgs{
		Service:  service,
		Endpoint: endpoint,
	}
	data, err := json.Marshal(args)
	if err != nil {
		panic(err)
	}
	transport.Call(server, "Put", data)

}

func Get(service string) string {
	args := api.GetArgs{
		Service: service,
	}
	data, err := json.Marshal(args)
	if err != nil {
		panic(err)
	}
	reponse := transport.Call(server, "Get", data)
	var res string
	err = json.Unmarshal(reponse, &res)
	if err != nil {
		panic(err)
	}
	return res
}

var server string
