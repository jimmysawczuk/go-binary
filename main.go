package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"go/format"
	"io"
	"os"
	"text/template"
)

var package_name string = "main"
var function_name string = ""
var debug bool = true
var output string = ""

var compiled_template *template.Template

const (
	NewLine string = "\n"
)

func init() {
	flag.StringVar(&package_name, "p", "main", "The package name to use")
	flag.StringVar(&function_name, "f", "getData", "The function name to use")
	flag.BoolVar(&debug, "debug", false, "Output the raw data as a comment in the output")
	flag.StringVar(&output, "out", "", "File to which to output the binary")

	compiled_template = template.Must(template.New("full").Parse(`package {{ .PackageName }}

import(
	"compress/gzip"
	"bytes"
	"io"
)

func {{ .FunctionName }}() ([]byte, error) {
{{ .DebugData }}

	in := bytes.NewBuffer([]byte{{ .CompressedData }})
	out := bytes.NewBuffer([]byte{})

	gz, err := gzip.NewReader(in)
	if err != nil {
		return []byte{}, err
	}
	io.Copy(out, gz)

	return out.Bytes(), nil
}
`))
}

type Data struct {
	Raw []byte
}

func main() {
	flag.Parse()

	fp := os.Stdin
	in := new(bytes.Buffer)
	_, err := in.ReadFrom(fp)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		data := Data{in.Bytes()}

		debug_data := ""
		if debug {
			debug_data = `/* RAW INPUT DATA:` + NewLine +
				data.String() + NewLine +
				`*/`
		}

		out := new(bytes.Buffer)

		err = compiled_template.Execute(out, struct {
			PackageName    string
			FunctionName   string
			DebugData      string
			CompressedData string
		}{
			PackageName:    package_name,
			FunctionName:   function_name,
			CompressedData: data.CompressedString(),
			DebugData:      debug_data,
		})

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		fmtted, _ := format.Source(out.Bytes())
		if output != "" {
			fp, err := os.OpenFile(output, os.O_TRUNC+os.O_CREATE+os.O_WRONLY, 0644)
			if err == nil {
				io.Copy(fp, bytes.NewBuffer(fmtted))
				fp.Close()
			} else {
				fmt.Fprintln(os.Stderr, err.Error())
			}
		} else {
			io.Copy(os.Stdout, bytes.NewBuffer(fmtted))
		}
	}
}

func (d Data) Compressed() []byte {
	compressed := new(bytes.Buffer)
	gz, _ := gzip.NewWriterLevel(compressed, gzip.BestCompression)
	io.Copy(gz, bytes.NewBuffer(d.Raw))
	gz.Close()

	return compressed.Bytes()
}

func (d Data) CompressedString() string {
	compressed := d.Compressed()

	str := `{`

	i := 0
	for _, v := range compressed {
		if i%12 == 0 {
			str = str + NewLine
			i = 0
		}

		str = str + fmt.Sprintf("0x%02x,", v)
		i++
	}

	str = str + `}`

	return str
}

func (d Data) String() string {
	return bytes.NewBuffer(d.Raw).String()
}

func (d Data) Len() int {
	return len(d.Raw)
}
