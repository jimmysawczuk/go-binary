# go-binary
[![Go Report Card](https://goreportcard.com/badge/github.com/jimmysawczuk/go-binary)](https://goreportcard.com/report/github.com/jimmysawczuk/go-binary)

Takes data from `os.Stdin` and transforms it into a Go source code file suitable for including in your project.

## Example

```bash
# Install go-binary
$ go get github.com/jimmysawczuk/go-binary

# Download some test content from the web
$ curl http://www.nytimes.com > index.html

# Pipe the content into go-binary to generate a .go file with our content in it
$ cat index.html | go-binary -f="getData" -out="$GOPATH/src/github.com/jimmysawczuk/go-binary/go-binary-example/get_data.go"

# Install the example, which depends on our file
$ go install github.com/jimmysawczuk/go-binary/go-binary-example

# Run the example, which just prints the content to the screen
$ go-binary-example

# Print the generated .go file to see what it looks like
$ cat $GOPATH/src/github.com/jimmysawczuk/go-binary/go-binary-example/get_data.go
```

## License

go-binary is released under [the MIT license][license].

  [license]: https://github.com/jimmysawczuk/go-binary/blob/master/LICENSE
