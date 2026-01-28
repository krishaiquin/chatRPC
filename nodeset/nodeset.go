package nodeset

import (
	nodesetManager "chatRPC/lib/nodesetManager/rpc/clientStub"
	"chatRPC/nodeset/api"
	"slices"
	"sync"
)

func Add(addr string, username string) api.AddRet {
	mx.Lock()
	id := nextId
	nextId += 1
	node := api.Node{
		NodeId:   id,
		UserName: username,
		Addr:     addr,
	}
	cluster = append(cluster, node)
	nodeset := api.AddRet{
		NodeId:  id,
		NodeSet: cluster,
	}
	mx.Unlock()

	go requestAdd(node)
	return nodeset
}

func Delete(nodeId uint32) {
	for index, n := range cluster {
		if nodeId == n.NodeId {
			mx.Lock()
			cluster = slices.Delete(cluster, index, index+1)
			mx.Unlock()
			break
		}
	}

	go reqeuestDelete(nodeId)
}

func requestAdd(node api.Node) {
	for _, destinationNode := range cluster {
		if destinationNode.NodeId == node.NodeId {
			continue
		}
		nodesetManager.AddMember(destinationNode.Addr, node)
	}
}

func reqeuestDelete(nodeId uint32) {
	for _, destinationNode := range cluster {
		nodesetManager.RemoveMember(destinationNode.Addr, nodeId)
	}
}

func init() {
	nextId = 0
}

var nextId uint32
var cluster []api.Node
var mx sync.Mutex
