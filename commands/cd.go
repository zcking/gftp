package commands

import (
	"github.com/zcking/gftp/client"
)

// CdCommand is a command for changing directory on the server
type CdCommand struct {
	raw string
}

// Execute sends the "cd" command to the server
func (c *CdCommand) Execute(cli client.Client) error {
  return cli.RunString(c.raw)
}
