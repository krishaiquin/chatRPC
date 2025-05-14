package nodeset

import "fmt"

func Add(addr string) {
	cluster = append(cluster, addr)
	fmt.Printf("Nodes: %v\n", cluster)
}

// func GetNodes() string {
// 	return fmt.Sprintf("Nodeset: %v\n", cluster)
// }

/**
* Maybe add delete node when they exit from the nodeset
 */

var cluster = make([]string, 0, 1)
