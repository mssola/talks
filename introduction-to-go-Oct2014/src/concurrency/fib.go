package main

import (
	"fmt"
)

//START2 OMIT
func fib(n int) chan int {
	c := make(chan int)

	go func() {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			a, b = b, a+b
			c <- a
		}
		close(c)
	}()
	return c
}

//STOP2 OMIT

// START1 OMIT
func main() {
	for i := range fib(10) {
		fmt.Println(i)
	}
}

// STOP1 OMIT
