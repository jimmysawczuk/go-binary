#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o build/go-binary .
docker build -t jimmysawczuk/go-binary .
docker push jimmysawczuk/go-binary
rm -rf build
