package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

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

//START1 OMIT
func main() {
	// Let there be channels.
	bob := speak("Bob")
	alice := speak("Alice")
	timeout := time.After(time.Duration(2) * time.Second) // HL

	// And let them talk.
	for {
		select {
		case s := <-bob:
			fmt.Println(s)
		case s := <-alice:
			fmt.Println(s)
		case <-timeout: // HL
			fmt.Printf("Timed out!\n") // HL
			os.Exit(0)                 // HL
		}
	}
}

//STOP1 OMIT
