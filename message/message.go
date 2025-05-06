package message

import "fmt"

func Echo(message string) string {
	return fmt.Sprintf("Received message: %s", message)
}
