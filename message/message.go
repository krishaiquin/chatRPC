package message

import "fmt"

func Send(message string) string {
	return fmt.Sprintf("Received message: %s", message)
}
