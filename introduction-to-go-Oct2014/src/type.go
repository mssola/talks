package main

import "fmt"

// START1 OMIT
type User struct {
	Name, Lastname string
	Age            int `json:"-"`
}

// STOP1 OMIT

// START3 OMIT
func (u User) Greet() {
	fmt.Printf("Hello! I'm %v\n", u.Name)
}

//STOP3 OMIT

// START2 OMIT
type MyOwnString string

//STOP2 OMIT

func main() {
}
