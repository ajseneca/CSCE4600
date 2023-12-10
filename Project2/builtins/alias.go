package builtins

import (
	"fmt"
	"io"
	"strings"
	"os"
	"bufio"
	"path/filepath"
)

func AssignAlias(w io.Writer, args ...string) error {
	HomeDir, _ = os.UserHomeDir()
	path := filepath.Join(HomeDir, "CSCE4600", "Project2", "aliases.txt")
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)

	if (err != nil) {
		return fmt.Errorf("Could not open/create \"aliases.txt\". Error: %w", err)
	}

	defer file.Close()

	if (len(args) < 1) {
		return fmt.Errorf("%w", ErrInvalidSyntax)
	}
	
	tempString := strings.Join(args, "")
	firstChar := '('
	lastChar := ')'

	
	if (!strings.HasPrefix(tempString, string(firstChar))) {
		return fmt.Errorf("%w", ErrInvalidSyntax)
	}

	if (!strings.HasSuffix(tempString, string(lastChar))) {
		return fmt.Errorf("%w", ErrInvalidSyntax)
	}
	
	str := strings.Trim(tempString, "(")
	aliasStr := strings.Trim(str, ")")
	aliasStr += ";"
	var commandResult string
	commandPart := strings.Split(aliasStr, "=")
	commandResult = commandPart[0]
	var aliasList []string
	aliasScanner := bufio.NewScanner(file)

	for aliasScanner.Scan() {
		aliasList = append(aliasList, aliasScanner.Text())
	}

	file.Close()

	for i := 0; i < len(aliasList); i++ {
		if (!strings.HasSuffix(aliasList[i], ";")) {
			aliasList[i] += ";\n"
		}
		if (strings.HasSuffix(aliasList[i], ";")) {
			aliasList[i] += "\n"
		}
	}

	finalList := append(aliasList, aliasStr)
	finalStr := strings.Join(finalList, "")
	cmdExist := CheckCommandExists(commandResult)

	if (!cmdExist) {
		return fmt.Errorf("Command not found. Alias cannot be created.")
	}

	overwrite, _ := os.Create(path)

	_, writeErr2 := overwrite.WriteString(finalStr)

	overwrite.Close()

	if writeErr2 != nil {
		return fmt.Errorf("Could not write alias to \"aliases.txt\". %w", writeErr2)
	}

	fmt.Fprintln(w, "Alias saved successfully!")

	return nil
}

func CheckCommandExists(check string) bool {
	switch check {
	case "cd":
		return true
	case "env":
		return true
	case "pwd":
		return true
	case "time":
		return true
	case "echo":
		return true
	case "alias":
		return true
	case "history":
		return true
	case "exit":
		return true
	}

	return false
}

func CheckForAlias(w io.Writer, args ...string) (string, bool) {
	HomeDir, _ = os.UserHomeDir()
	path := filepath.Join(HomeDir, "CSCE4600", "Project2", "aliases.txt")
	file, err := os.Open(path)

	if err != nil {
		fmt.Fprintln("Could not open \"aliases.txt\". Error: %w", err)
		return "", false
	}

	defer file.Close()

	buffer := make([]byte, 1024)
	var aliases []string

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		contents := strings.TrimSpace(string(buffer[:n]))
		aliases = append(aliases, contents)
	}

	var argsStr string = strings.Join(args, "")
	var aliasListSplit []string = strings.Split(strings.Join(aliases, ";"), ";")

	for _, alias := range aliasListSplit {		
		commandPart := strings.Split(alias, "=")
		if len(commandPart) > 1 {
			commandResult := strings.Trim(commandPart[1], "\n")

			if (commandResult == argsStr) {
				return strings.Trim(commandPart[0], "\n"), true
			}
		}
	}

	return "", false
}