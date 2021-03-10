package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var fileFlag = flag.String("file", "", "file as input")

func main() {
	flag.Parse()
	output := calcMD5(input())
	fmt.Printf("%x", output)
}

func input() io.Reader {
	if *fileFlag != "" {
		f, err := os.Open("file.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		return f
	} else {
		args := os.Args
		if len(args) != 1 {
			fmt.Println(args[0])
		}
		return strings.NewReader(args[1])
	}
}

func calcMD5(reader io.Reader) []byte {
	h := md5.New()
	io.Copy(h, reader)
	return h.Sum(nil)
}
