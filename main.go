package main

import "fmt"

func main() {
	s := "qwerty test one two tung"
	result := beisen(s)
	fmt.Println(result)

	result = avocado(s)
	fmt.Println(result)

	result2 := beisen2(4) // 1 + 2 + 3 + 4 = 10
	fmt.Println(result2)

	result2 = beisen3(4)
	fmt.Println(result2)
}
