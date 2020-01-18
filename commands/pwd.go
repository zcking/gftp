package commands

import (
	"github.com/zcking/gftp/client"
)

// PwdCommand is a command for printing the current working directory
// on the remote machine
type PwdCommand struct {}

// Execute sends the "pwd" command to the server
func (c *PwdCommand) Execute(cli client.Client) error {
  return cli.RunString("pwd")
}
