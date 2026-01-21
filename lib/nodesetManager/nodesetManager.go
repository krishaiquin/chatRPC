package nodesetManager

import (
	"chatRPC/dlog"
	"chatRPC/lib/transport"
	"chatRPC/nodeset/api"
	nodeset "chatRPC/nodeset/rpc/clientStub"
	"slices"
	"sync"
)

type Cluster struct {
	NodeId   uint32 //this node's id
	NodeSet  []api.Node
	OnChange func([]api.Node)
	mx       sync.Mutex
}

func CreateCluster(username string) {
	cluster.mx.Lock()
	cluster.NodeId, cluster.NodeSet = nodeset.Add(transport.GetAddress(), username)
	dlog.Printf("nodeId: %d\n", cluster.NodeId)
	cluster.mx.Unlock()
}

func AddMember(node api.Node) {
	cluster.mx.Lock()
	cluster.NodeSet = append(cluster.NodeSet, node)
	cluster.mx.Unlock()

	if cluster.OnChange != nil {
		cluster.OnChange(cluster.NodeSet)
	}
}

func RemoveMember(nodeId uint32) {
	for index, n := range cluster.NodeSet {
		if nodeId == n.NodeId {
			cluster.mx.Lock()
			cluster.NodeSet = slices.Delete(cluster.NodeSet, index, index+1)
			cluster.mx.Unlock()
			break
		}
	}
}

func GetId() uint32 {
	return cluster.NodeId
}

func GetNode(nodeId uint32) api.Node {
	for _, node := range cluster.NodeSet {
		if node.NodeId == nodeId {
			return node
		}
	}

	return api.Node{}
}

func GetNodeSet() []api.Node {
	return cluster.NodeSet
}

func GetCluster() *Cluster {
	return cluster
}

func init() {
	cluster = &Cluster{}

}

var cluster *Cluster
