package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 3
	}()

	n := <-c
	fmt.Printf("%v\n", n)
}
