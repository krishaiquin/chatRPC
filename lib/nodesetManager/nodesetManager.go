package nodesetManager

import (
	"chatRPC/lib/transport"
	"chatRPC/nodeset/api"
	nodeset "chatRPC/nodeset/rpc/clientStub"
	"slices"
	"sync"
)

type Cluster struct {
	NodeId   uint32 //this node's id
	NodeSet  []api.Node
	OnChange func(*DiffCluster)
	mx       sync.Mutex
}

type DiffCluster struct {
	AddedNodes   []api.Node
	RemovedNodes []api.Node
}

func CreateCluster(username string) {
	cluster.mx.Lock()
	cluster.NodeId, cluster.NodeSet = nodeset.Add(transport.GetAddress(), username)
	cluster.mx.Unlock()
}

func AddMember(node api.Node) {
	cluster.mx.Lock()
	cluster.NodeSet = append(cluster.NodeSet, node)
	diff.AddedNodes = append(diff.AddedNodes, node)
	cluster.mx.Unlock()

	if cluster.OnChange != nil {
		cluster.OnChange(diff)
	}
}

func RemoveMember(nodeId uint32) {
	for index, n := range cluster.NodeSet {
		if nodeId == n.NodeId {
			cluster.mx.Lock()
			cluster.NodeSet = slices.Delete(cluster.NodeSet, index, index+1)
			diff.RemovedNodes = append(diff.RemovedNodes, n)
			cluster.mx.Unlock()

			if cluster.OnChange != nil {
				cluster.OnChange(diff)
			}

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
	diff = &DiffCluster{}
}

var cluster *Cluster
var diff *DiffCluster
