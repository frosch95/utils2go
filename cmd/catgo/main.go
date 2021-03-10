package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Provide at least one file")
	}
	for _, file := range os.Args[1:] {
		printFile(file)
	}
}

func printFile(file string) {
	fd, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	io.Copy(os.Stdout, fd)
}
