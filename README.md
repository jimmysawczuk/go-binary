# go-binary
Takes data from `os.Stdin` and transforms it into a Go source code file suitable for including in your project.

## Example

```bash
$ curl http://www.google.com > index.html && \
    go install github.com/jimmysawczuk/go-binary && \
    cat index.html | go-binary -f="getData" -out="$GOPATH/src/github.com/jimmysawczuk/go-binary/example/get_data.go" && \
    go install github.com/jimmysawczuk/go-binary/example && \
    example
```
