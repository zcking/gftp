package shell

import (
	"bufio"
	"fmt"
	"strings"
)

const (
	prompt = "gftp> "
)

// ReadLine reads a line of input (input until a newline) from a reader
func ReadLine(r *bufio.Reader) (string, error) {
	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// PrintPrompt outputs the gftp prompt for input
func PrintPrompt() {
	fmt.Print(prompt)
}

// Newline prints a newline feed to stdout
func Newline() {
	fmt.Printf("\n")
}
