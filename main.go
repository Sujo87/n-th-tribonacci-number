package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter n:")
	fmt.Scan(&n)

	fmt.Println(tribonacciFuncSpaceOptimized(n))

}

func tribonacciFuncSpaceOptimized(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}

	if n <= 3 {
		return 1
	}
	a := 0
	b := 1
	c := 1
	d := a + b + c

	i := 2
	for i < n {
		d = a + b + c
		a = b
		b = c
		c = d

		i++
	}

	return d
}
