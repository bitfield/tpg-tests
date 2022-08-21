package greet

import "fmt"

func Greet() {
	fmt.Print("Your name? ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s!\n", name)
}
