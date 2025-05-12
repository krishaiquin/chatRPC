package nodeset

import "fmt"

func Add(addr string) string {
	cluster = append(cluster, addr)
	fmt.Printf("Nodes: %v\n", cluster)
	return "Success!"
}

// func GetNodes() string {
// 	return fmt.Sprintf("Nodeset: %v\n", cluster)
// }

var cluster = make([]string, 0, 1)

/**
* Maybe add delete node when they exit from the nodeset
 */
