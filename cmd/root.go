package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/zcking/gftp/client"
	"github.com/zcking/gftp/commands"

	"github.com/spf13/cobra"
	"github.com/zcking/gftp/shell"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gftp",
		Short: "gFTP is a simple FTP and SFTP client",
		Long:  "gFTP is a simple FTP and SFTP client written in Go",
		Run:   run,
		Args:  cobra.MinimumNArgs(1),
	}
)

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
}

func run(cmd *cobra.Command, args []string) {
	target, err := client.NewDestination(args[len(args)-1])
	if err != nil {
		exit(err)
	}

	// Establish connection with Client
	sftp := &client.SFTP{}
	if err = sftp.Connect(target); err != nil {
		exit(err)
	}

	reader := bufio.NewReader(os.Stdin)

	for sftp.IsConnected() {
		shell.PrintPrompt()
		rawInput, err := shell.ReadLine(reader)
		if err != nil {
			exit(err)
		}
		comm, err := commands.ParseCommand(rawInput)
		if err != nil {
			shell.Print(err)
		} else {
			// Execute the parsed command
			if comm != nil {
				comm.Execute(sftp)
			}
		}
	}
}

func exit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
