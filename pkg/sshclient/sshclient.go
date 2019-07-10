package sshclient

import (
	"io/ioutil"
	"net"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// TerminalConn is use for making ssh connection with pty request
func TerminalConn(user string, keyPath string, ipAddr string, port string) {
	// Joining ip address and port as a strings
	value := []string{}
	value = append(value, ipAddr)
	value = append(value, port)
	ipPort := strings.Join(value, ":")

	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey([]byte(key))
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	conn, err := ssh.Dial("tcp", ipPort, config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
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

	fileDescriptor := int(os.Stdin.Fd())

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
}