package main

import (
	"log"
	"os"
)

func main() {
	err := os.Truncate("test.txt", 12)
	if err != nil {
		log.Fatal(err)
	}
}
