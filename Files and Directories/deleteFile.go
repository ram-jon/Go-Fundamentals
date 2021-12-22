package main

import (
	"log"
	"os"
)

func main() {
	err := os.RemoveAll("copy")
	if err != nil {
		log.Fatal(err)
	}

}
