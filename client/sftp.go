package client

import (
	"fmt"
  "io"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

// SFTP is a connection wrapper for a SFTP network connection
type SFTP struct {
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
func (sftp *SFTP) Connect(target *Destination) error {
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
	sftp.conn = connection
	sftp.isConnected = true
	fmt.Printf("Connected to %s.\n", target.Host)

  // Start a new SSH session with the remote machine
  // which will be used for all command execution
  sftp.session, err = sftp.newSession()
  if err != nil {
    return err
  }

  // Get STDIN handle to be able to pass commands to remote
  sftp.stdin, err = sftp.session.StdinPipe()
  if err != nil {
    return err
  }

  sftp.session.Stdout = os.Stdout

  // Shell() starts a login shell on the remote machine
	return sftp.session.Shell()
}

func (sftp *SFTP) newSession() (*ssh.Session, error) {
	return sftp.conn.NewSession()
}

// RunString writes a string payload to the server
func (sftp *SFTP) RunString(payload string) error {
  // send the command to the remote machine
  sftp.stdin.Write([]byte(payload + "\n"))
  return nil
}

// Close disconnects the SFTP connection
func (sftp *SFTP) Close() error {
	sftp.isConnected = false
  if err := sftp.stdin.Close(); err != nil {
    return err
  }
  sftp.session.Close()      // close the SSH sesssion
	return sftp.conn.Close()  // close the TCP connection
}

// IsConnected returns whether or not the SFTP
// connection is established
func (sftp *SFTP) IsConnected() bool {
	return sftp.isConnected
}
