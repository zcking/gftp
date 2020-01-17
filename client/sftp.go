package client

import (
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh/knownhosts"

	"golang.org/x/crypto/ssh"
)

// SFTP is a connection wrapper for a SFTP network connection
type SFTP struct {
	conn *ssh.Client
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

	sshConfig := &ssh.ClientConfig{
		User: target.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(target.Pass), // TODO: parameterize
		},
		HostKeyCallback: hostKeyCallback,
	}

	connection, err := ssh.Dial("tcp", target.String(), sshConfig)
	if err != nil {
		return err
	}
	sftp.conn = connection

	return nil
}

func (sftp *SFTP) session() (*ssh.Session, error) {
	return sftp.conn.NewSession()
}

// RunString writes a string payload to the server
func (sftp *SFTP) RunString(payload string) ([]byte, error) {
	sh, err := sftp.session()
	if err != nil {
		return nil, err
	}
	return sh.Output(payload)
}

// Close disconnects the SFTP connection
func (sftp *SFTP) Close() error {
	return sftp.conn.Close()
}
