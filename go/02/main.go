package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Print("What is the input string? ")
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimRight(s, "\n")
	n := len(s)
	fmt.Printf("%v has %v characters.\n", s, n)

}
