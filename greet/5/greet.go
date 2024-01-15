package greet

import (
	"bufio"
	"fmt"
	"io"
)

func Greet(in io.Reader, out io.Writer) {
	fmt.Fprint(out, "Your name? ")
	scanner := bufio.NewScanner(in)
	if !scanner.Scan() {
		return
	}
	fmt.Fprintf(out, "Hello, %s!\n", scanner.Text())
}
