package main

import "fmt"

func IsSorted(f func(a int, b int) bool, array []int) bool {
	for i := 0; i < len(array)-1; i++ {
		if !f(array[i], array[i+1]) {
			return false
		}
	}
	return true
}

func f(a, b int) bool {
	if b <= a {
		return false
	}
	return true
}

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Print(result1)
	fmt.Print("\n")
	fmt.Print(result2)
}
