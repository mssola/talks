package main

//START4 OMIT
type Person struct {
	name, lastname string
	age            int
}

var p1, p2 *Person

var persons []Person

var i int

//STOP4 OMIT

//START1 OMIT
func function(f func(x int) int) func() int {
	return func() int {
		return f(2) + 2
	}
}

//STOP1 OMIT

//START2 OMIT
func foo(s string) (value int, err error)

//STOP2 OMIT

//START3 OMIT
func main() {
	a := 0 // "a" declared as an integer
}

//STOP3 OMIT
