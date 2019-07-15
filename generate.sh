#!/bin/bash
set -vx

# 64bit builder
env GOOS=darwin GOARCH=amd64 go build -o bin/64bit/darwin/gossh cmd/gossh/main.go
env GOOS=linux GOARCH=amd64 go build -o bin/64bit/linux/gossh cmd/gossh/main.go
env GOOS=windows GOARCH=amd64 go build -o bin/64bit/window/gossh cmd/gossh/main.go
