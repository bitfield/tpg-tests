package greet

import (
	"bufio"
	"fmt"
	"os"
)

func Main() int {
	fmt.Println("Your name? ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return 1
	}
	fmt.Printf("Hello, %s!\n", scanner.Text())
	return 0
}
