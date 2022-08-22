package greet7

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
	name := scanner.Text()
	fmt.Println("Your favourite food? ")
	if !scanner.Scan() {
		return 1
	}
	food := scanner.Text()
	fmt.Printf("Hello, %s. Care to join me for some %s?\n", name,
		food)
	return 0
}
