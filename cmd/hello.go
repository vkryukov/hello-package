package main

import (
	"bufio"
	"fmt"
	"github.com/vkryukov/hello-package"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What's you name? ")
	scanner.Scan()
	hello.Greeter(scanner.Text())
	hello.RunDetector(scanner)
}
