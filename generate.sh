#!/bin/bash
set -vx

# 64bit builder
env GOOS=darwin GOARCH=amd64 go build -o ~/gosshbin/64bit/darwin/gossh_darwin_amd64 cmd/gossh/main.go
shasum ~/gosshbin/64bit/darwin/gossh_darwin_amd64 | awk '{print $1}' > ~/gosshbin/64bit/darwin/gossh_darwin_amd64.shasum
env GOOS=linux GOARCH=amd64 go build -o ~/gosshbin/64bit/linux/gossh_linux_amd64 cmd/gossh/main.go
shasum  ~/gosshbin/64bit/linux/gossh_linux_amd64 | awk '{print $1}' >  ~/gosshbin/64bit/linux/gossh_linux_amd64.shasum
# env GOOS=windows GOARCH=amd64 go build -o ~/gosshbin/64bit/window/gossh cmd/gossh/main.go
