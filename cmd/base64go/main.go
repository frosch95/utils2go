package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var fileFlag = flag.String("file", "", "file as input")
var modeFlag = flag.String("mode", "encode", "encode or decode as mode")

func main() {
	flag.Parse()

	if *modeFlag != "encode" && *modeFlag != "decode" {
		log.Fatal("mode must be encode or decode")
	}

	input, err := input()
	if err != nil {
		log.Fatal(err)
	}

	inputAsString := streamToString(input)

	var output string
	switch {
	case *modeFlag == "decode":
		output = decode(inputAsString)
	default:
		output = encode(inputAsString)
	}

	fmt.Println(output)
}

func streamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
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
	case len(flag.Args()) == 1:
		return strings.NewReader(flag.Args()[0]), nil
	case len(flag.Args()) == 0:
		return os.Stdin, nil
	default:
		return nil, errors.New("Provide a file or an argument or use the standard input!")
	}
}

func encode(msg string) string {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}

func decode(msg string) string {
	decoded, _ := base64.StdEncoding.DecodeString(msg)
	return string(decoded)
}
