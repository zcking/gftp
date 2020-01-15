package client

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh/knownhosts"

	"golang.org/x/crypto/ssh"
)

// SFTP is a connection wrapper for a SFTP network connection
type SFTP struct {
	conn    *ssh.Client
	session *ssh.Session
	stdout  io.Reader
	stdin   io.WriteCloser
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
	hostKeyCallback, err := getHostKeyCallback(target.host)
	if err != nil {
		return err
	}

	sshConfig := &ssh.ClientConfig{
		User: os.Getenv("GFTP_USER"), // TODO: parameterize
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("GFTP_PASS")), // TODO: parameterize
		},
		HostKeyCallback: hostKeyCallback,
	}

	connection, err := ssh.Dial("tcp", target.String(), sshConfig)
	if err != nil {
		return err
	}
	sftp.conn = connection
	session, err := connection.NewSession()
	if err != nil {
		return err
	}
	sftp.session = session

	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}
	sftp.stdout = stdout
	stdin, err := session.StdinPipe()
	if err != nil {
		return err
	}
	sftp.stdin = stdin

	return nil
}

// SendString writes a string payload to the server
func (sftp *SFTP) SendString(payload string) error {
	return sftp.session.Run(payload)
}

// RecvString receives a string from the server stdout
func (sftp *SFTP) RecvString() (string, error) {
	bytes, err := ioutil.ReadAll(sftp.stdout)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Close disconnects the SFTP connection
func (sftp *SFTP) Close() error {
	if err := sftp.stdin.Close(); err != nil {
		return err
	}
	if err := sftp.session.Close(); err != nil {
		return err
	}
	return sftp.conn.Close()
}
