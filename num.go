package main

import "fmt"

func num() {
	var n int
	fmt.Scanln(&n)
	if n > 0 {
		fmt.Println("positive")
	} else if n < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("zero")
	}
}
