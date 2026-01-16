package message

import (
	"chatRPC/nodeset/api"
	"fmt"
)

func Send(node api.Node, message string) {
	fmt.Printf("%s: %s", node.UserName, message)
}
