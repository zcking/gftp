package client

import (
	"fmt"
	"os/user"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// Client is a FTP or SFTP wrapper
type Client interface {
	Connect(*Destination) error
	RunString(string) error
	Close() error
	IsConnected() bool
}

// Destination is the target server
type Destination struct {
	Host string
	Port int
	User string
	Pass string
	Path string
}

// parseHostAndPath parses a string
// from format host[:path] into the separate
// (host, path) values
func parseHostAndPath(s string) (string, string) {
	toks := strings.SplitN(s, ":", 2)

	if len(toks) > 1 {
		return toks[0], toks[1]
	}

	return s, "./"
}

// inputPassword prompt the username@host's password
// and returns the input as a string
func inputPassword(username string, host string) (string, error) {
	fmt.Printf("%s@%s's password: ", username, host)
	bytePassword, err := terminal.ReadPassword(syscall.Stdin)
	fmt.Printf("\n")
	if err != nil {
		return "", err
	}
	return string(bytePassword), nil
}

// NewDestination parses the required argument of gFTP
// which must be [user@]host[:path]
func NewDestination(destStr string) (*Destination, error) {
	var serverUser string
	var host string
	var path string
	var password string

	toks := strings.SplitN(destStr, "@", 2)
	if len(toks) > 1 {
		serverUser = toks[0]
		host, path = parseHostAndPath(toks[1])
	} else {
		// default user to the current username on host machine
		curUser, err := user.Current()
		if err != nil {
			return nil, err
		}
		serverUser = curUser.Username
		host, path = parseHostAndPath(destStr)
	}

	// Prompt user for password
	password, err := inputPassword(serverUser, host)
	if err != nil {
		return nil, err
	}

	return &Destination{
		User: serverUser,
		Pass: password,
		Host: host,
		Port: 22, // TODO: make configurable
		Path: path,
	}, nil
}

func (d *Destination) String() string {
	return fmt.Sprintf("%v:%v", d.Host, d.Port)
}
