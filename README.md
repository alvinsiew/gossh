# Gossh

![](https://img.shields.io/github/release/alvinsiew/gossh.svgsvg?style=flat)

Command line SSH client.

```text
About
Support linux and mac. (Window is not working 100% currently, will try to make it workable in next version)
Host informations etc(ip address, port, key) will be encrypted on rest in datafile.
Does not require to install ssh to use Gossh.
No installation require. Just need to copy binary to client machine.
```

## Getting Started

Copy from gossh/bin to your machine /bin folder for linux, and gossh is ready to be use.

To compile on your own. Please choose your own favour of os to compile.

```bash
Example:
Macos
env GOOS=darwin GOARCH=amd64 go build -o bin/64bit/darwin/gossh cmd/gossh/main.go
Linux
env GOOS=linux GOARCH=amd64 go build -o bin/64bit/linux/gossh cmd/gossh/main.go
Window
env GOOS=windows GOARCH=amd64 go build -o bin/64bit/window/gossh cmd/gossh/main.go
```

Usage:

```golang

Usage of ./gossh:
  -add
        Add host:
        Usage: gossh -add -host <hostname|mandatory> -ip <ip address|mandatory> -user <userid|non-mandatory> -port <ssh port|non-mandatory> -key <private key|non-mandatory>
  -c    Connection to server:
        Usage: gossh -conn <hostname>
  -del
        Hostname to delete
  -host string
        Hostname
  -ip string
        Adding or changing IP address for host
  -key string
        Setup key to for server connection. Using default key if not specific. (default "nokey")
  -l    List all hosts config
         -l info to list more infor
         -l key to list private key
  -pass string
        User password
  -port string
        Port Number (default "22")
  -user string
        User (default "alvinsiew")
exit status 2

$ go run cmd/gossh/main.go -h
Usage of /var/folders/33/3_dzcxkn2wg2zvkk_l4977fc0000gn/T/go-build358610656/b001/exe/main:
  -add
        Add host:
        Usage: gossh -add -host <hostname|mandatory> -ip <ip address|mandatory> -user <userid|non-mandatory> -port <ssh port|non-mandatory> -key <private key|non-mandatory>
  -c    Connection to server:
        Usage: gossh -conn <hostname>
  -del
        Hostname to delete
  -host string
        Hostname
  -ip string
        Adding or changing IP address for host
  -key string
        Setup key to for server connection. Using default key if not specific. (default "nokey")
  -l    List all hosts config
         -l info
        to list more infor
         -l key
        to list private key
  -pass string
        User password
  -port string
        Port Number (default "22")
  -user string
        User (default "default user")
```

Example:

```golang
To add a host:
gossh -add -host server-test -ip 192.168.1.23 -user centos -port 22 -key /home/hello/id_rsa

To remove a host:
gossh -del server-test

To list all hosts:
gossh -l

To list more info on hosts:
gossh -l info

To list host private key:
gossh -l key

To connect to host:
gossh -c server-test
```

### Prerequisites

No prerequistes required.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
