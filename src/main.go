package main

import (
	"log"
	"time"
)

func main() {
	log.Print("APPSTATUS started.")

	for {
		log.Print(".")
		time.Sleep(time.Second)
	}
}
