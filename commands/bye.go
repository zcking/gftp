package commands

import (
	"github.com/zcking/gftp/client"
)

// ByeCommand is a command for listing files on the server
type ByeCommand struct{}

// Execute disconnects from the server
func (c *ByeCommand) Execute(cli client.Client) error {
	cli.Close()
	return nil
}
