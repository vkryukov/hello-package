package main

import (
	"bufio"
	"fmt"
	"github.com/endeveit/guesslanguage"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What's you name? ")
	scanner.Scan()
	Greeter(scanner.Text())

	for {
		fmt.Print("Enter some text: ")
		scanner.Scan()
		language, err := guesslanguage.Guess(scanner.Text())
		if err != nil {
			fmt.Printf("Error detecting language: %s\n", err)
		} else {
			fmt.Printf("Detected language: %s\n", language)
		}
	}
}
