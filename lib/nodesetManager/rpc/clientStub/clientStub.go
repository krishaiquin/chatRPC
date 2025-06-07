package clientStub

import (
	"chatRPC/lib/transport"
	"chatRPC/nodeset/api"
	"encoding/json"
	"log"
)

func Update(destination string, cluster []api.Node) {

	//marshal the args
	data, err := json.Marshal(cluster)
	if err != nil {
		panic(err)
	}

	//send it to destination
	response := transport.Call(destination, "Update", data)
	if response != nil {
		log.Fatalf("Error occured: response is not nil.")
	}

}
