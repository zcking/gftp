package commands

import (
	"github.com/zcking/gftp/client"
)

// LsCommand is a command for listing files on the server
type LsCommand struct {
	raw string
}

// Execute sends the "ls" command to the server
// and receives the listing resopnse
func (c *LsCommand) Execute(cli client.Client) error {
  return cli.RunString(c.raw)
}
