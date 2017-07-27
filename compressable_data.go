package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

type compressableData []byte

func (d compressableData) compressed() []byte {
	if compression <= 0 {
		return []byte(d)
	}
	compressed := new(bytes.Buffer)
	gz, _ := gzip.NewWriterLevel(compressed, compression)
	io.Copy(gz, bytes.NewBuffer(d))
	gz.Close()

	return compressed.Bytes()
}

func (d compressableData) compressedString() string {
	compressed := d.compressed()

	str := `{`

	i := 0
	for _, v := range compressed {
		if i%12 == 0 {
			str = str + newline
			i = 0
		}

		str = str + fmt.Sprintf("0x%02x,", v)
		i++
	}

	str = str + `}`

	return str
}

func (d compressableData) String() string {
	return string([]byte(d))
}

func (d compressableData) safeString() string {
	s := d.String()
	s = strings.Replace(s, "/*", "/ *", -1)
	s = strings.Replace(s, "*/", "* /", -1)
	return s
}

func (d compressableData) Len() int {
	return len(d)
}
