package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	var c int
	a, b := 0, 1

	fib := func(no int) int {
		c = a + b
		a = b
		b = c
		return a
	}
	return fib
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
