package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	div := a / b
	mod := a % b
	return div, mod
}

func printRange(a int, b int) {
	for i := a; i <= b; i++ {
		fmt.Println(i)
	}
}
