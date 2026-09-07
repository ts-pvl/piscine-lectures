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

// первое слово
func beisen(s string) string {
	result := ""
	for _, r := range s {
		if r == ' ' {
			break
		}
		// result += string(r)
		result = result + string(r)
	}
	return result
}

func avocado(s string) string {
	result := ""
	secondWord := false
	for _, r := range s {
		if r == ' ' {
			if secondWord == false {
				secondWord = true
			} else {
				break
			}
		}
		if secondWord {
			result = result + string(r)
		}
	}
	return result
}

func retainHalf(s string) string {
	var res string
	for i := 0; i < len(s)/2; i++ {
		res = res + string(s[i])
	}
	return res
}
