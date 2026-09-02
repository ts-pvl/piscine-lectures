package main

import (
	"fmt"
)

func test2() {
	var age int
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("child")
	}
}
