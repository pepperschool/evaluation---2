package main

import "fmt"

func AppendRange(min, max int) []int {
	var x []int
	if min < max {
		for i := min; i < max+1; i++ {
			x = append(x, i)
		}
	} else {
		return nil
	}
	return x
}

func main() {
	a1 := AppendRange(5, 10)
	a2 := AppendRange(10, 5)

	fmt.Print(a1)
	fmt.Print(a2)
}
