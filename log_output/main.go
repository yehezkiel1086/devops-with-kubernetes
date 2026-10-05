package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	// generate random UUID and store to memory
	storedString := uuid.New().String()
	fmt.Printf("Initialized with random string: %s\n\n", storedString)

	// setup timer for every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// output the stored random UUID every 5 seconds
	for {
		select {
		case t := <-ticker.C:
			// format to match millisecond precision and UTC (Z)
			timestamp := t.UTC().Format("2006-01-02T15:04:05.000Z")
			fmt.Printf("%s: %s\n", timestamp, storedString)
		}
	}
}