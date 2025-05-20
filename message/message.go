package message

import (
	nodeset "chatRPC/nodeset/rpc/clientStub"
	"fmt"
)

func Send(message string) string {
	// reqeuest for list of nodes
	cluster := nodeset.GetNodes()
	fmt.Printf("Nodes: %v\n", cluster)
	// send the message to all nodes in the cluster
	// for _, str := range cluster {
	// 	fmt.Println(str)
	// }
	return fmt.Sprintf("Received message: %s", message)
}
