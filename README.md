# go-binary
Takes data from `os.Stdin` and transforms it into a Go source code file suitable for including in your project.

## Example

```bash
$ curl http://www.google.com > index.html && \
	go install go-binary && go install go-binary/example && \
	cat index.html | go-binary -f="getData" -out="example/get_data.go" && \
	example
```
