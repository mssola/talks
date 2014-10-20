package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//START1 OMIT
func fanIn(first, second chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-first:
				c <- s
			case s := <-second:
				c <- s
			}
		}
	}()
	return c
}

//STOP1 OMIT

func speak(msg string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

//START2 OMIT
func main() {
	// Let there be channels.
	bob := speak("Bob")
	alice := speak("Alice")
	c := fanIn(bob, alice) // HL
	timeout := time.After(time.Duration(2) * time.Second)

	// And let them talk.
	for {
		select {
		case s := <-c: // HL
			fmt.Println(s) // HL
		case <-timeout:
			fmt.Printf("Timed out!\n")
			os.Exit(0)
		}
	}
}

//STOP2 OMIT
