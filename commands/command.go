package commands

import (
	"fmt"
	"strings"

	"github.com/zcking/gftp/client"
)

// Command is a executable FTP statement
type Command interface {
	Execute(cli client.Client) error
}

// ParseCommand parses a raw user input into
// a gFTP command, if valid
func ParseCommand(raw string) (Command, error) {
	if raw == "" {
		return nil, fmt.Errorf("")
	}

	tokens := strings.SplitN(raw, " ", 2)
	comm := strings.ToLower(tokens[0])

	switch comm {
	case "ls":
		return &LsCommand{raw: raw}, nil
  case "cd":
    return &CdCommand{raw: raw}, nil
  case "pwd":
    return &PwdCommand{}, nil
	case "bye":
		return &ByeCommand{}, nil
	default:
		return nil, fmt.Errorf("unsupported command '%v'\n", comm)
	}
}
