package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		if i == 0 {
			// Name of the executable
			fmt.Printf("Executable: %s\n", arg)
		} else {
			Greeter(arg)
		}
	}
}
