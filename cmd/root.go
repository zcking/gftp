package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/zcking/gftp/shell"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gftp",
		Short: "gFTP is a simple FTP and SFTP client",
		Long:  "gFTP is a simple FTP and SFTP client written in Go",
		Run:   run,
	}
)

type connection struct {
	host string
	port int
}

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
}

func run(cmd *cobra.Command, args []string) {
	conn, err := parseHostAndPort(args)
	if err != nil {
		exit(err)
	}

	fmt.Fprintf(os.Stdout, "connecting to %s:%d ...\n", conn.host, conn.port) // TODO: establish connection over TCP

	reader := bufio.NewReader(os.Stdin)

	for {
		shell.PrintPrompt()
		comm, err := shell.ReadLine(reader)
		if err != nil {
			exit(err)
		}
		fmt.Printf("'%v' command not implemented\n", comm)
		shell.Newline()
	}
}

func exit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

// parseHostAndPort parses the last two arguments
// [host [port]] as the host and port values.
// default connection is nil
func parseHostAndPort(args []string) (*connection, error) {
	numArgs := len(args)

	if numArgs < 1 {
		return nil, nil
	} else if numArgs >= 2 {
		host := args[numArgs-2]
		port, err := strconv.ParseInt(args[numArgs-1], 10, 32)
		if err != nil {
			return nil, err
		}

		return &connection{
			host,
			(int)(port),
		}, nil
	} else {
		return &connection{
			host: args[numArgs-1],
			port: 0,
		}, nil
	}
}
