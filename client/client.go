package client

import (
	"fmt"
	"strconv"
)

// Client is a FTP or SFTP wrapper
type Client interface {
	Connect(*Destination) error
	SendString(string) error
	RecvString() (string, error)
	Close() error
}

// Destination is the target server
type Destination struct {
	host string
	port int
}

// NewDestination parses the last two arguments
// [host [port]] as the host and port values.
// default connection is nil
func NewDestination(args []string) (*Destination, error) {
	numArgs := len(args)

	if numArgs < 1 {
		return nil, nil
	} else if numArgs >= 2 {
		host := args[numArgs-2]
		port, err := strconv.ParseInt(args[numArgs-1], 10, 32)
		if err != nil {
			return nil, err
		}

		return &Destination{
			host,
			(int)(port),
		}, nil
	} else {
		return &Destination{
			host: args[numArgs-1],
			port: 0,
		}, nil
	}
}

func (d *Destination) String() string {
	return fmt.Sprintf("%v:%v", d.host, d.port)
}
