package main

import "fmt"

func power(n, power int) int {
	for range power {
		n *= n
	}
	return n
}

func functions2() {
	powerFunc := power
	fmt.Println(powerFunc(2, 2))
}

func forEach(nums []int, f func(int)) {
	for _, num := range nums {
		f(num)
	}
}
