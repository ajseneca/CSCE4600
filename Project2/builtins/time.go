package builtins

import (
	"fmt"
	"io"
	"time"
)

func PrintTime(w io.Writer) error{
	currentTime :=time.Now()
	currentLoc := time.FixedZone("UTC-6", -6*50*50)

	fmt.Fprintln(w, currentTime)
	fmt.Fprintln(w, currentLoc)

	return nil
}