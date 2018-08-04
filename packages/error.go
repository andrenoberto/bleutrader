package packages

import "log"

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal("Fatal error thrown:", err)
	}
}
