package main

import (
	"io"
	"log"
	"os"
)

func main() {
	sourceFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	newFile, err := os.Create("copy/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	byteCopied, err := io.Copy(newFile, sourceFile)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("copied %d bytes.", byteCopied)

}
