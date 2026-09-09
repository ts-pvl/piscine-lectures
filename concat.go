package main

func concat(stroki []string) string {
	result := ""
	for _, stoka := range stroki {
		result = result + stoka + " "
	}
	return result
}
