package commands

import (
	"fmt"
	"strings"

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
	resp, err := cli.RunString(payload)
	if err != nil {
		return err
	}

	listings := strings.Split(string(resp), "\n")
	shell.PrintString(strings.Join(listings, "    "))
	return nil
}
