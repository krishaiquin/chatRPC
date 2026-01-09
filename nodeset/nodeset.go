package nodeset

import (
	"chatRPC/dlog"
	nodesetManager "chatRPC/lib/nodesetManager/rpc/clientStub"
	"chatRPC/nodeset/api"
	"slices"
	"sync"
)

func Add(addr string) uint32 {
	mx.Lock()
	id := nextId
	nextId += 1
	cluster = append(cluster, api.Node{NodeId: id, Addr: addr})
	mx.Unlock()
	//send ack to requester

	//notify all nodes in cluster
	//fmt.Printf("Nodeset server cluster: %v\n", cluster)
	go notify()

	return id
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
	go notify()

}

func notify() {

	for _, n := range cluster {
		nodesetManager.Update(n.Addr, cluster)
		dlog.Printf("node %d - %s updated!", n.NodeId, n.Addr)
	}

}

func init() {
	nextId = 0
}

var nextId uint32
var cluster []api.Node
var mx sync.Mutex
