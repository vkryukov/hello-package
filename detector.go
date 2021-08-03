package hello

import (
	"bufio"
	"fmt"
	"github.com/endeveit/guesslanguage"
)

func RunDetector(scanner *bufio.Scanner) {
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
