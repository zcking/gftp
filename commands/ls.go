package commands

import (
	"fmt"

	"github.com/zcking/gftp/shell"

	"github.com/zcking/gftp/client"
)

// LsCommand is a command for listing files on the server
type LsCommand struct {
	path string
}

// Execute sends the "ls" command to the server
// and receives the listing resopnse
func (c *LsCommand) Execute(cli client.Client) error {
	payload := fmt.Sprintf("ls %v\n", c.path)
	cli.SendString(payload)
	resp, err := cli.RecvString()
	if err != nil {
		return err
	}

	shell.PrintString(resp)
	return nil
}
