package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	var res interface{}
	var err error
	res, err = getData()
	if err != nil {
		fmt.Fprintln(os.Stderr, "getData failed:", err.Error())
		os.Exit(2)
	}

	switch ty := res.(type) {
	case []byte:
		fmt.Print(string(fromBytes(ty)))
	case io.Reader:
		fmt.Print(string(fromReader(ty)))
	}
}

func fromBytes(in []byte) []byte {
	return in
}

func fromReader(in io.Reader) []byte {
	by, _ := ioutil.ReadAll(in)
	return by
}
