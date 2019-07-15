# Gossh&#33;
Command line SSH client.

Features:
Host informations etc(ip address, port, key) will be encrypted on rest in datafile.
Does not require to install ssh to use Gossh!.
No installation require. Just need to copy binary to client machine.

## Getting Started
Copy from gossh/bin to your machine /bin folder for linux, and gossh is ready to be use.

To compile on your own. Please choose your own favour of os to compile.
Example:
env GOOS=darwin GOARCH=amd64 go build -o bin/64bit/darwin/gossh cmd/gossh/main.go

### Prerequisites

No prerequistes required.


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details