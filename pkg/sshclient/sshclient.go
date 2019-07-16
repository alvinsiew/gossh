package sshclient

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"runtime"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// TerminalConn is use for making ssh connection with pty request
func TerminalConn(user string, keyPath []byte, ipAddr string, port string, pass string) error {
	// Joining ip address and port as a strings
	value := []string{}
	value = append(value, ipAddr)
	value = append(value, port)
	ipPort := strings.Join(value, ":")
	var config *ssh.ClientConfig

	config, err := ClientConfig(user, keyPath, pass)
	if err != nil {
		panic("Failed to pass client config: " + err.Error())
	}

	conn, err := ssh.Dial("tcp", ipPort, config)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer conn.Close()

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := conn.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	// Set IO
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	var fileDescriptor int
	
	if runtime.GOOS == "windows" {
		fileDescriptor = int(os.Stdout.Fd())
	} else {
		fileDescriptor = int(os.Stdin.Fd())
	}

	// terminal connected to the given file descriptor into raw mode and returns the previous state of the terminal so that it can be restored.
	if terminal.IsTerminal(fileDescriptor) {
		originalState, err := terminal.MakeRaw(fileDescriptor)
		if err != nil {
			log.Fatalf("Connect terminal to file descriptor in raw mode failed: %s", err)
		}
		defer terminal.Restore(fileDescriptor, originalState)

		termWidth, termHeight, err := terminal.GetSize(fileDescriptor)
		if err != nil {
			log.Fatalf("Getting terminal size failed: %s", err)
		}

		err = session.RequestPty("xterm-256color", termHeight, termWidth, modes)
		if err != nil {
			log.Fatalf("Request Pty failed: %s", err)
		}
	}

	// Starts a login shell on the remote host
	err = session.Shell()
	if err != nil {
		log.Fatalf("Starts a login shell failed: %s", err)
	}
	session.Wait()

	return err
}

// KeyToBytes convert key to bytes
func KeyToBytes(keyPath string) []byte {
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic(err)
	}
	return key
}

// ClientConfig ssh login authentication method
func ClientConfig(user string, keyPath []byte, pass string) (*ssh.ClientConfig, error) {
	var signer ssh.Signer
	var err error
	var config *ssh.ClientConfig
	keyLen := len(keyPath)

	if keyLen != 0 {
		signer, err = ssh.ParsePrivateKey([]byte(keyPath))
		if err != nil {
			panic(err)
		}
	}

	if pass == "" && keyLen <= 0 {
		config = &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.KeyboardInteractive(challenge),
			},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		}
	} else if keyLen > 0 {
		config = &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		}
	} else if pass != "" {
		config = &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password(pass),
			},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		}
	}

	return config, err
}

// Interaction between server and client
func challenge(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
	answers = make([]string, len(questions))
	for n, q := range questions {
		fmt.Printf("Enter %s\n", q)
		bytePassword, err := terminal.ReadPassword(0)
		if err != nil {
			panic(err)
		}
		password := string(bytePassword)

		answers[n] = password
	}

	return answers, nil
}
