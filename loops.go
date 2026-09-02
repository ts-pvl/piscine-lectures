package main

import "fmt"

func test3() {
	for i := 0; i <= 10; i = i + 1 {
		fmt.Println(i)
	}
}

func evens() {
	// number = number + 2
	for number := 0; number <= 100; number += 2 {
		fmt.Println(number)
	}
}

func odds() {
	for number := 0; number <= 100; number++ {
		// bool only
		if number%2 == 1 {
			fmt.Println(number)
		}
	}
}

func while() {
	var n int
	for true {
		fmt.Scanln(&n)
		if n > 0 {
			fmt.Println("positive")
		} else if n == 0 {
			fmt.Println("zero")
		} else {
			fmt.Println("negative")
		}
	}
}
