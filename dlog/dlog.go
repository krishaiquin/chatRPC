package dlog

/**
*	Debug log
*
*	Prints debug logs when debug is enable.
*	Exmaple: DEBUG="1" go run chat.go <dbEndpoit>
*
 */

import (
	"log"
	"os"
)

func Printf(format string, v ...any) {
	if debug {
		log.Printf(format, v...)
	}
}

var debug = os.Getenv("DEBUG") == "1"
