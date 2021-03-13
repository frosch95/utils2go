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
	errorHandling(err)
	output := calcMD5(input)
	fmt.Printf("%x", output)
}

func input() (io.Reader, error) {
	switch {
	case *fileFlag != "":
		f, err := os.Open(*fileFlag)
		errorHandling(err)
		defer func() {
			err = f.Close()
		}()
		errorHandling(err)
		return f, nil
	case len(flag.Args()) == 1:
		return strings.NewReader(flag.Args()[0]), nil
	case len(flag.Args()) == 0:
		return os.Stdin, nil
	default:
		return nil, errors.New("provide a file or an argument or use the standard input")
	}
}

func calcMD5(reader io.Reader) []byte {
	h := md5.New()
	_, err := io.Copy(h, reader)
	errorHandling(err)
	return h.Sum(nil)
}

func errorHandling(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
