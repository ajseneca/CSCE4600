package builtins

import (
	"fmt"
	"io"
	"strings"
	"os"
	"bufio"
	"path/filepath"
)

func PrintHistory(w io.Writer) error {
	HomeDir, _ = os.UserHomeDir()
	path := filepath.Join(HomeDir, "CSCE4600", "Project2", "history.txt")
	file, err := os.OpenFile(path, os.O_RDONLY, 0777)

	if (err != nil) {
		return fmt.Errorf("Could not open/create history.txt")
	}

	defer file.Close()

	historyScanner := bufio.NewScanner(file)
	var history []string

	for historyScanner.Scan() {
		history = append(history, string(historyScanner.Text() + "\n"))
	}

	history[len(history) - 1] = strings.Trim(history[len(history) - 1], "\n")
	historyStr := strings.Join(history, "")

	fmt.Fprintln(w, "History:")
	fmt.Fprintln(w, historyStr)	

	return nil
}

func AddHistory(w io.Writer, name string) error {
	HomeDir, _ = os.UserHomeDir()
	path := filepath.Join(HomeDir, "CSCE4600", "Project2", "history.txt")
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if (err != nil) {
		return fmt.Errorf("Could not open \"history.txt\". Error: %w", err)
	}

	defer file.Close()
	line := name + "\n"
	_, writeErr := file.WriteString(line)

	if writeErr != nil {
		return fmt.Errorf("Could not write history to \"history.txt\". %w", writeErr)
	}

	return nil
}

func DeleteHistory(w io.Writer) error {
	HomeDir, _ = os.UserHomeDir()
	path := filepath.Join(HomeDir, "CSCE4600", "Project2", "history.txt")
	
	err := os.Remove(path)
	if (err != nil) {
		return fmt.Errorf("Could not delete \"history.txt\". %w", err)
	}

	return nil
}