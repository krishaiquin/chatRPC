package nodesetManager

import (
	"chatRPC/dlog"
	"chatRPC/lib/transport"
	"chatRPC/nodeset/api"
	nodeset "chatRPC/nodeset/rpc/clientStub"
	"sync"
)

type Cluster struct {
	NodeId  uint32 //this node's id
	NodeSet []api.Node
	mx      sync.Mutex
}

// Only called by client nodes
func CreateCluster() {
	cluster.mx.Lock()
	cluster.NodeId = nodeset.Add(transport.GetAddress())
	cluster.mx.Unlock()
}

// calls by nodeset services
func Update(nodeset []api.Node) {
	dlog.Printf("updating my local copy of cluster")
	cluster.mx.Lock()
	cluster.NodeSet = make([]api.Node, len(nodeset))
	copy(cluster.NodeSet, nodeset)
	cluster.mx.Unlock()
	dlog.Printf("Cluster: ")
	for _, node := range cluster.NodeSet {
		dlog.Printf("%s, ", node.Addr)
	}
}

func GetId() uint32 {
	return cluster.NodeId
}

func GetCluster() []api.Node {
	return cluster.NodeSet
}

func init() {
	cluster = &Cluster{}

}

var cluster *Cluster
