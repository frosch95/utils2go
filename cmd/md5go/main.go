package main

import (
	"crypto/md5"
	"errors"
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
	input, err := input()
	if err != nil {
		log.Fatal(err)
	}
	output := calcMD5(input)
	fmt.Printf("%x", output)
}

func input() (io.Reader, error) {
	switch {
	case *fileFlag != "":
		f, err := os.Open(*fileFlag)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		return f, nil
	case len(os.Args) == 2:
		return strings.NewReader(os.Args[1]), nil
	case len(os.Args) == 1:
		return os.Stdin, nil
	default:
		return nil, errors.New("Provide a file or an argument or use the standard input!")
	}
}

func calcMD5(reader io.Reader) []byte {
	h := md5.New()
	io.Copy(h, reader)
	return h.Sum(nil)
}
