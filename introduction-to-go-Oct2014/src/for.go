package main

func main() {
	ary := []int{1, 2, 3, 4, 5}

	// START1 OMIT
	for i := 0; i < len(ary); i++ {
		// Do something.
	}
	// STOP1 OMIT

	// START2 OMIT
	i := 0
	for i < len(ary) {
		// Do something.
		i++
	}
	// STOP2 OMIT

	// START3 OMIT
	for key, value := range ary {
		// Do something with the key and the value.
	}
	// STOP3 OMIT
}
