package nodeset

import (
	nodesetManager "chatRPC/lib/nodesetManager/rpc/clientStub"
	"chatRPC/nodeset/api"
	"sync"
)

func Add(addr string) uint32 {
	mx.Lock()
	id := nextId
	nextId += 1
	cluster = append(cluster, api.Node{NodeId: id, Addr: addr})
	mx.Unlock()
	notify()

	return id
}

func notify() {
	for _, n := range cluster {
		nodesetManager.Update(n.Addr, cluster)
	}
}

func init() {
	nextId = 0

}

var nextId uint32
var cluster []api.Node
var mx sync.Mutex
