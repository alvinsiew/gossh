# Gossh

Command line SSH client.

```text
Features:
Host informations etc(ip address, port, key) will be encrypted on rest in datafile.
Does not require to install ssh to use Gossh!.
No installation require. Just need to copy binary to client machine.
```

## Getting Started

Copy from gossh/bin to your machine /bin folder for linux, and gossh is ready to be use.

To compile on your own. Please choose your own favour of os to compile.
Example:
env GOOS=darwin GOARCH=amd64 go build -o bin/64bit/darwin/gossh cmd/gossh/main.go

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
         -l all to list all values
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

To list all hosts with all informations:
gossh -l all

To connect to host:
gossh -c server-test
```

### Prerequisites

No prerequistes required.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
