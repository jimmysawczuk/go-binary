package main

import (
	"bytes"
	"compress/gzip"
	"strings"

	"github.com/pkg/errors"
)

type payload []byte

func compress(in payload, level int) (payload, error) {
	if level <= 0 {
		return in, nil
	}

	tmp := &bytes.Buffer{}
	gz, err := gzip.NewWriterLevel(tmp, level)
	if err != nil {
		return nil, errors.Wrap(err, "gzip: new writer")
	}

	if _, err := gz.Write(in); err != nil {
		return nil, errors.Wrap(err, "gzip: write")
	}

	if err := gz.Close(); err != nil {
		return nil, errors.Wrap(err, "gzip: close")
	}

	return payload(tmp.Bytes()), nil
}

func (d payload) SafeString() string {
	s := string([]byte(d))
	s = strings.Replace(s, "/*", "/ *", -1)
	s = strings.Replace(s, "*/", "* /", -1)
	return s
}

func (d payload) FormatForCode() string {
	return formatForCode([]byte(d))
}

func (d payload) Len() int {
	return len(d)
}
