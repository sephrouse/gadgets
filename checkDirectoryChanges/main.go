package main

import (
	"log"

	"github.com/rjeczalik/notify"
)

func main() {
	c := make(chan notify.EventInfo, 1)

	// Set up a watchpoint listening on events within current working directory.
	// Dispatch each create and remove events separately to c.
	if err := notify.Watch(".", c, notify.Create, notify.Remove); err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)

	// Block until an event is received.
	for {
		ei := <-c
		log.Println("Got event:", ei)
	}

}
