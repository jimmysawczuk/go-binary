package main

import (
	"bytes"
	"fmt"
)

func main() {
	data, _ := getData()
	fmt.Println(bytes.NewBuffer(data).String())
}
