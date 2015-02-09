# go-binary
Takes data from `os.Stdin` and transforms it into a Go source code file suitable for including in your project.

## Example

```bash
$ go get github.com/jimmysawczuk/go-binary && \
    curl http://www.google.com > index.html && \
    cat index.html | go-binary -f="getData" -out="$GOPATH/src/github.com/jimmysawczuk/go-binary/go-binary-example/get_data.go" && \
    go install github.com/jimmysawczuk/go-binary/go-binary-example && \
    go-binary-example
```

You can then do a `cat $GOPATH/src/github.com/jimmysawczuk/go-binary/go-binary-example/get_data.go` to see the output that was automatically generated and compiled into `go-binary-example`.
