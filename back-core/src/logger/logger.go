package logger

import "log"

func Log(result string, handler string) {
	log.Printf("-----------------\n")
	log.Printf("handler: %v\n", handler)
	log.Printf("result: %v\n", result)
}
