package main

import (
	"fmt"
	"os"
)

func main() {
	word1 := os.Args[1]
	word2 := os.Args[2]
	fmt.Print(word1 + word2)
}
