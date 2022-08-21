package greet4

import (
	"fmt"
	"io"
)

func Greet(in io.Reader, out io.Writer) {
	fmt.Fprint(out, "Your name? ")
	var name string
	fmt.Fscanln(in, &name)
	fmt.Fprintf(out, "Hello, %s!\n", name)
}
