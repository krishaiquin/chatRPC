package nodeset

import (
	nodesetManager "chatRPC/lib/nodesetManager/rpc/clientStub"
	"chatRPC/nodeset/api"
	"log"
	"sync"
)

func Add(addr string) uint32 {
	mx.Lock()
	id := nextId
	nextId += 1
	cluster = append(cluster, api.Node{NodeId: id, Addr: addr})
	mx.Unlock()
	//fmt.Printf("Nodeset server cluster: %v\n", cluster)
	go notify()

	return id
}

func notify() {

	for _, n := range cluster {
		nodesetManager.Update(n.Addr, cluster)
		log.Printf("node %d - %s updated!", n.NodeId, n.Addr)
	}

}

func init() {
	nextId = 0

}

var nextId uint32
var cluster []api.Node
var mx sync.Mutex
