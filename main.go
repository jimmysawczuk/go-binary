package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/pkg/errors"
)

var packageName = "main"
var functionName = ""
var debug = true
var output = ""
var compression = 9
var useReaders = false

const (
	newline string = "\n"
)

func init() {
	flag.StringVar(&packageName, "p", "main", "The package name to use")
	flag.StringVar(&functionName, "f", "getData", "The function name to use")
	flag.BoolVar(&debug, "debug", false, "Output the raw data as a comment in the output")
	flag.StringVar(&output, "out", "", "File to which to output the binary")
	flag.IntVar(&compression, "compression", 9, "gzip compression level to use (0 to 9, 0 for no compression)")
	flag.BoolVar(&useReaders, "readers", false, "Use io.Readers in function signatures instead of []byte.")
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
	by, err := ioutil.ReadAll(fp)
	if err != nil {
		fatalErr(errors.Wrap(err, "read"))
	}

	data := payload(by)
	compressed, err := compress(data, compression)
	if err != nil {
		fatalErr(errors.Wrap(err, "compress"))
	}

	out := &bytes.Buffer{}
	vars := exportData{
		PackageName:  packageName,
		FunctionName: functionName,

		Data:        data,
		Compressed:  compressed,
		Compression: compression,
		PctSavings:  100 - 100*float64(compressed.Len())/float64(data.Len()),

		UseReaders: useReaders,
		Debug:      debug,
		ParseTime:  time.Now(),
	}
	if err := exportTemplate.Execute(out, vars); err != nil {
		fatalErr(errors.Wrap(err, "execute template"))
	}

	fmtted, err := format.Source(out.Bytes())
	if err != nil {
		fatalErr(errors.Wrap(err, "gofmt"))
	}

	var outputWriter io.WriteCloser = os.Stdout
	if output != "" {
		outputWriter, err = os.OpenFile(output, os.O_TRUNC+os.O_CREATE+os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}
	}

	io.Copy(outputWriter, bytes.NewReader(fmtted))
	outputWriter.Close()
}

func fatalErr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
