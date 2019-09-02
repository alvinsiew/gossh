# Gossh

![Version](https://img.shields.io/github/release/alvinsiew/gossh.svg?style=flat)

Gossh is a command line SSH client where configurations for private key, password, IP address, port and user will be stored in a database file and encrypted.

**Note:**<br/>
Supports Linux and MacOS. (Windows is not working 100% currently, will try to make it workable in next version)

* Host information (IP address, port, key, etc) will be encrypted at rest in a database file.
* Installation of SSH is not required to use Gossh.
* No installation required. Just copy the binary to client machine.

## Install

### From GitHub Releases

Please see [GitHub Releases](https://github.com/alvinsiew/gossh/releases).<br/>
Available binaries are:
* MacOS
* Linux

Copy to the /bin folder of your machine for Linux or /usr/local/bin folder for MacOS, and gossh is ready to use.

### go get

`$ go get -u github.com/alvinsiew/gossh/cmd/gossh`

### Self-compile

```bash
$ git clone https://github.com/alvinsiew/gossh.git

# MacOS
$ env GOOS=darwin GOARCH=amd64 go build -o bin/64bit/darwin/gossh cmd/gossh/main.go

# Linux
$ env GOOS=linux GOARCH=amd64 go build -o bin/64bit/linux/gossh cmd/gossh/main.go

# Window
$ env GOOS=windows GOARCH=amd64 go build -o bin/64bit/window/gossh cmd/gossh/main.go
```

## Usage

```bash
$ gossh -h
Usage of gossh:
  -add
      Add host:
      Usage: gossh -add -host <hostname|mandatory> -ip <ip address|mandatory> -user <userid|non-mandatory> -port <ssh port|non-mandatory> -key <private key|non-mandatory>
  -c
      Connection to server:
      Usage: gossh -c <hostname>
  -del
    	Hostname to delete
  -host string
    	Hostname
  -ip string
    	Adding or changing IP address for host
  -key string
    	Setup key to for server connection. Using default key if not specific. (default "nokey")
  -l	List all hosts config
    	 -l info
    	to list more infor
    	 -l key
    	to list private key
  -pass string
    	User password
  -port string
    	Port Number (default "22")
  -user string
    	User (default "alvinsiew")
```

## Example

```bash
# To add a host
gossh -add -host server-test -ip 192.168.1.23 -user centos -port 22 -key /home/hello/id_rsa

# To remove a host
gossh -del server-test

# To list all hosts
gossh -l

# To list more info on hosts
gossh -l info

# To list host private key
gossh -l key

#To connect to host
gossh -c server-test
```

### Prerequisites

No prerequisites required.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
