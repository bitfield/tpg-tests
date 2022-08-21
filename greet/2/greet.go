package greet2

import (
	"fmt"
	"io"
)

func Greet(out io.Writer) {
	fmt.Fprint(out, "Your name? ")
	var name string
	fmt.Scanln(&name)
	fmt.Fprintf(out, "Hello, %s!\n", name)
}
