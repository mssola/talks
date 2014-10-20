package main

import (
	"fmt"
)

//START3 OMIT
func filter(in chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			if v%2 == 0 {
				out <- v
			}
		}
		close(out)
	}()
	return out
}

//STOP3 OMIT

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
	for i := range filter(fib(10)) { // HL
		fmt.Println(i)
	}
}

// STOP1 OMIT
