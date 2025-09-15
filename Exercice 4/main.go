package main

import "fmt"

func Countif(tab []string) int {
	count := 0
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) {
			count++
		}
	}
	return count
}

func f(s string) bool {
	if s == " " {
		return false
	}
	return true
}

func main() {
	array := []string{"les test", " ", " ", "LES TEST", " "}
	fmt.Print(Countif(array))
}
