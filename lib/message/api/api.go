package api

import "chatRPC/nodeset/api"

type SendArgs struct {
	From    api.Node
	Message string
}
