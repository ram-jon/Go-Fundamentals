package main

import (
	"log"
	"os"
)

func main() {
	oldLocation := "testing.txt"
	newLocation := "test/testing.txt"
	err := os.Rename(oldLocation, newLocation)

	if err != nil {
		log.Fatal(err)
	}
}
