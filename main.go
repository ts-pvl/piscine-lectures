package main

import "fmt"

func printNumber(n int) {
	fmt.Println(n)
}

func main() {
	forEach([]int{1, 2, 3}, printNumber)
}
