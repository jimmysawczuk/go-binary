package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io"
	"os"
	"time"
)

var packageName = "main"
var functionName = ""
var debug = true
var output = ""
var compression = 9

const (
	newline string = "\n"
)

func init() {
	flag.StringVar(&packageName, "p", "main", "The package name to use")
	flag.StringVar(&functionName, "f", "getData", "The function name to use")
	flag.BoolVar(&debug, "debug", false, "Output the raw data as a comment in the output")
	flag.StringVar(&output, "out", "", "File to which to output the binary")
	flag.IntVar(&compression, "compression", 9, "gzip compression level to use (0 to 9, 0 for no compression)")
}

func main() {
	flag.Parse()

	if compression < 0 {
		compression = 0
	}
	if compression > 9 {
		compression = 9
	}

	fp := os.Stdin
	in := new(bytes.Buffer)
	_, err := in.ReadFrom(fp)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	data := compressableData(in.Bytes())

	debugData := ""
	if debug {
		debugData = `/* RAW INPUT DATA:` + newline +
			data.safeString() + newline +
			`*/`
	}

	out := new(bytes.Buffer)

	err = exportTemplate.Execute(out, exportData{
		PackageName:    packageName,
		FunctionName:   functionName,
		CompressedData: data.compressedString(),
		DebugData:      debugData,
		ParseDate:      time.Now(),
		Compression:    compression,
		PctSavings:     100 - 100*float64(len(data.compressed()))/float64(data.Len()),
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	fmtted, err := format.Source(out.Bytes())
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	var outputWriter io.WriteCloser
	if output != "" {
		outputWriter, err = os.OpenFile(output, os.O_TRUNC+os.O_CREATE+os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}
	} else {
		outputWriter = os.Stdout
	}

	io.Copy(outputWriter, bytes.NewBuffer(fmtted))
	outputWriter.Close()
}
