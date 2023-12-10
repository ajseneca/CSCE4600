package builtins

import (
	"fmt"
	"os"
	"io"
)

func PresentWorkingDirectory(w io.Writer) error {
	currentDir, err := os.Getwd()
	fmt.Fprintln(w, currentDir)
	return err
}