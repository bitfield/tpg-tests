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
	name := scanner.Text()
	fmt.Println("Your favourite food? ")
	if !scanner.Scan() {
		os.Exit(1)
	}
	food := scanner.Text()
	fmt.Printf("Hello, %s. Care to join me for some %s?\n", name, food)
}
