package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/ajseneca/CSCE4600/Project2/builtins"
)

func main() {
	exit := make(chan struct{}, 2) // buffer this so there's no deadlock.
	runLoop(os.Stdin, os.Stdout, os.Stderr, exit)
}

func runLoop(r io.Reader, w, errW io.Writer, exit chan struct{}) {
	var (
		input    string
		err      error
		readLoop = bufio.NewReader(r)
	)
	for {
		select {
		case <-exit:
			_, _ = fmt.Fprintln(w, "exiting gracefully...")
			return
		default:
			if err := printPrompt(w); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if input, err = readLoop.ReadString('\n'); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if err = handleInput(w, input, exit); err != nil {
				_, _ = fmt.Fprintln(errW, err)
			}
		}
	}
}

func printPrompt(w io.Writer) error {
	// Get current user.
	// Don't prematurely memorize this because it might change due to `su`?
	u, err := user.Current()
	if err != nil {
		return err
	}
	// Get current working directory.
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// /home/User [Username] $
	_, err = fmt.Fprintf(w, "%v [%v] $ ", wd, u.Username)

	return err
}

func handleInput(w io.Writer, input string, exit chan<- struct{}) error {
	// Remove trailing spaces.
	input = strings.TrimSpace(input)

	// Split the input separate the command name and the command arguments.
	args := strings.Split(input, " ")
	name, args := args[0], args[1:]

	// Check for built-in commands.
	// New builtin commands should be added here. Eventually this should be refactored to its own func.

	var existState bool
	aliasName, existState := builtins.CheckForAlias(w, name)

	if (existState) {
		name = aliasName
	}

	switch name {
	case "cd":		
		err := builtins.AddHistory(w, name)
		return builtins.ChangeDirectory(args...)
	case "env":
		err := builtins.AddHistory(w, name)
		return builtins.EnvironmentVariables(w, args...)
	case "pwd":
		err := builtins.AddHistory(w, name)
		return builtins.PresentWorkingDirectory(w)
	case "time":
		err := builtins.AddHistory(w, name)
		return builtins.PrintTime(w)
	case "echo":
		err := builtins.AddHistory(w, name)
		return builtins.EchoText(w, args...)
	case "alias":
		err := builtins.AddHistory(w, name)
		return builtins.AssignAlias(w, args...)
	case "history":
		err := builtins.AddHistory(w, name)
		return builtins.PrintHistory(w)
	case "exit":
		err := builtins.DeleteHistory(w)
		exit <- struct{}{}
		return nil
	}

	if (err != nil) {
		fmt.Errorf("Error: ", err)
	}

	return executeCommand(name, args...)
}

func executeCommand(name string, arg ...string) error {
	// Otherwise prep the command
	cmd := exec.Command(name, arg...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}