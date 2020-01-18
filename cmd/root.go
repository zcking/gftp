package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/zcking/gftp/shell"
	"github.com/zcking/gsh/client"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gsh",
		Short: "gsh is a simple SSH client",
		Long:  "gsh is a simple SSH client written in Go",
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
	ssh := &client.GShell{}
	if err = ssh.Connect(target); err != nil {
		exit(err)
	}

	reader := bufio.NewReader(os.Stdin)

	for ssh.IsConnected() {
		shell.PrintPrompt()
		rawInput, err := shell.ReadLine(reader)
		if err != nil {
			exit(err)
		}

		if rawInput == "exit" {
			ssh.Close()
		} else {
			if err = ssh.RunString(rawInput); err != nil {
				exit(err)
			}
		}
	}
}

func exit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
