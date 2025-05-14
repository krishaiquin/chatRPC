package serverStub

import (
	"chatRPC/db"
	"chatRPC/db/api"
	"chatRPC/lib/transport"
	"encoding/json"
)

func Register() {
	transport.RegisterServerStub("Put", Put)
	transport.RegisterServerStub("Get", Get)
}

func Put(data []byte) []byte {
	//unmarshal
	var args api.PutArgs
	err := json.Unmarshal(data, &args)
	if err != nil {
		panic(err)
	}
	//call the procedure
	db.Put(args.Service, args.Endpoint)

	return nil
}

func Get(data []byte) []byte {
	var args api.GetArgs
	err := json.Unmarshal(data, &args)
	if err != nil {
		panic(err)
	}
	result := db.Get(args.Service)
	res, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return res

}
