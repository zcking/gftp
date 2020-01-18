package client

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

// GShell is a connection wrapper for a SSH network connection
type GShell struct {
	conn        *ssh.Client
	isConnected bool
	session     *ssh.Session
	stdin       io.WriteCloser
}

func getHostKeyCallback(host string) (ssh.HostKeyCallback, error) {
	knownHostFile := filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")
	callback, err := knownhosts.New(knownHostFile)
	if err != nil {
		return nil, err
	}
	return callback, err
}

// Connect establishes a TCP connection with a host
func (sh *GShell) Connect(target *Destination) error {
	hostKeyCallback, err := getHostKeyCallback(target.Host)
	if err != nil {
		return err
	}

	// Configure the SSH client for connecting
	// TODO: allow connecting w/ SSH key instead of password
	sshConfig := &ssh.ClientConfig{
		User: target.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(target.Pass),
		},
		HostKeyCallback: hostKeyCallback,
	}

	// Connect to the remote machine via SSH
	connection, err := ssh.Dial("tcp", target.String(), sshConfig)
	if err != nil {
		return err
	}
	sh.conn = connection
	sh.isConnected = true
	fmt.Printf("Connected to %s.\n", target.Host)

	// Start a new SSH session with the remote machine
	// which will be used for all command execution
	sh.session, err = sh.newSession()
	if err != nil {
		return err
	}

	// Get STDIN handle to be able to pass commands to remote
	sh.stdin, err = sh.session.StdinPipe()
	if err != nil {
		return err
	}

	sh.session.Stdout = os.Stdout

	// Shell() starts a login shell on the remote machine
	return sh.session.Shell()
}

func (sh *GShell) newSession() (*ssh.Session, error) {
	return sh.conn.NewSession()
}

// RunString writes a string payload to the server
func (sh *GShell) RunString(payload string) error {
	// send the command to the remote machine
	sh.stdin.Write([]byte(payload + "\n"))
	return nil
}

// Close disconnects the GShell connection
func (sh *GShell) Close() error {
	sh.isConnected = false
	if err := sh.stdin.Close(); err != nil {
		return err
	}
	sh.session.Close()     // close the SSH sesssion
	return sh.conn.Close() // close the TCP connection
}

// IsConnected returns whether or not the GShell
// connection is established
func (sh *GShell) IsConnected() bool {
	return sh.isConnected
}
