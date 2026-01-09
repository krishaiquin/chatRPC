package message

import (
	"fmt"
)

func Send(from uint32, message string) {
	fmt.Printf("Node %d: %s", from, message)
}
