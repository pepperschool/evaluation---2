package main

import "fmt"

func SplitwhiteSpaces(s string) []string {
	var x []string
	split := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			x = append(x, s[split:i])
			split = i
		}
	}
	return x
}

func main() {
	fmt.Print(SplitwhiteSpaces("LES TEST"))
}
