package commands

import (
	"github.com/zcking/gftp/client"
	"github.com/zcking/gftp/shell"
)

// LsCommand is a command for listing files on the server
type LsCommand struct {
	raw string
}

// Execute sends the "ls" command to the server
// and receives the listing resopnse
func (c *LsCommand) Execute(cli client.Client) error {
	resp, err := cli.RunString(c.raw)
	if err != nil {
		return err
	}

	shell.PrintString(string(resp))
	return nil
}
