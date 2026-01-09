package dlog

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
