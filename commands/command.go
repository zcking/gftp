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
	tokens := strings.Split(raw, " ")
	if len(tokens) == 0 || raw == "" {
		return nil, fmt.Errorf("")
	}

	comm := strings.ToLower(tokens[0])
	switch comm {
	case "ls":
		if len(tokens) >= 2 {
			return &LsCommand{path: tokens[1]}, nil
		}

		return &LsCommand{path: "./"}, nil

	default:
		return nil, fmt.Errorf("unsupported command '%v'\n", comm)
	}
}
