# Install go-binary
go install github.com/jimmysawczuk/go-binary/... && \

# Download some test content from the web, pipe it into go-binary.
curl -s 'https://baconipsum.com/api/?type=meat-and-filler&format=text' > test_in.txt

cat test_in.txt | go-binary -f="getData" -debug -out="$GOPATH/src/github.com/jimmysawczuk/go-binary/go-binary-example/get_data.go" && \

# Install the example, which depends on our file
go install github.com/jimmysawczuk/go-binary/go-binary-example && \

# Run the example, which just prints the content to the screen
go-binary-example > test_out.txt && \

# Compare raw input with go-binary-example output
md5 test_in.txt && md5 test_out.txt && \

# Cleanup
rm *.txt
