package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var iteration int
	var prev int
	var second_prev int
	return func() int {
		var fib int
		if iteration == 0 {
			fib = 0
		} else if iteration == 1 {
			fib = 1
		} else {
			fib = prev + second_prev
		}
		second_prev = prev
		prev = fib
		iteration += 1

		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
