package greet

import (
	"bufio"
	"fmt"
	"os"
)

func Main() {
	fmt.Println("Your name? ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		os.Exit(1)
	}
	fmt.Printf("Hello, %s!\n", scanner.Text())
}
