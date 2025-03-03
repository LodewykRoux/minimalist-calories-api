package errorHandling

import "log"

func FatalCheck(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
