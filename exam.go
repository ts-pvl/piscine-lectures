package main

func countChars(s string) int {
	return len(s)
}

func countLetters(s string) int {
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			count++
		}
	}
	return count
}

func perimeter(a int, b int) int {
	return 2*a + 2*b
}
