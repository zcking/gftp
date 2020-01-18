package shell

import (
	"bufio"
	"fmt"
	"strings"
)

const (
	prompt = "gsh# "
)

// ReadLine reads a line of input (input until a newline) from a reader
func ReadLine(r *bufio.Reader) (string, error) {
	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// PrintPrompt outputs the gsh prompt for input
func PrintPrompt() {
	fmt.Print(prompt)
}

// Newline prints a newline feed to stdout
func Newline() {
	fmt.Printf("\n")
}

// PrintString prints a string to the shell
func PrintString(s string) {
	fmt.Println(s)
}

// Print prints interfaces to the shell
func Print(a ...interface{}) {
	fmt.Print(a...)
}
