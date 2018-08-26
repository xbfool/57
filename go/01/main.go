package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	test, _ := reader.ReadString('\n')
	test = strings.TrimRight(test, "\n")
	fmt.Printf("Hello %s, nice to meet you!\n", test)
}
