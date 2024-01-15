package hello

import (
	"fmt"
	"os"
)

func Main() int {
	if len(os.Args[1:]) < 1 {
		fmt.Fprintln(os.Stderr, "usage: hello NAME")
		return 1
	}
	fmt.Println("Hello to you,", os.Args[1])
	return 0
}
