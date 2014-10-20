package main

import (
	"fmt"
	"math/rand"
	"time"
)

//START1 OMIT
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

//STOP1 OMIT

//START2 OMIT
func main() {
	// Let there be channels.
	bob := speak("Bob")
	alice := speak("Alice")

	// And let them talk.
	for {
		select {
		case s := <-bob:
			fmt.Println(s)
		case s := <-alice:
			fmt.Println(s)
		}
	}
}

//STOP2 OMIT
