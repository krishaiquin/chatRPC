package serverStub

import (
	"chatRPC/message"
	"encoding/json"
)

func Echo(args []byte) []byte {
	//unmarshal
	var msg string
	err := json.Unmarshal(args, &msg)
	if err != nil {
		panic(err)
	}
	//call the procedure
	res := message.Echo(msg)
	//marshall
	data, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	return data
}
