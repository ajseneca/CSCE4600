package builtins

import (
	"fmt"
	"io"
)

func EchoText(w io.Writer, args ...string) error{
	echoIn := args
	fmt.Fprintln(w, echoIn)

	return nil
}