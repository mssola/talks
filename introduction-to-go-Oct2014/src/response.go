package main

import (
	"encoding/json"
	"fmt"
)

//START1 OMIT
type Response struct {
	Message string `json:"msg,omitempty"`
	Error   string `json:"error,omitempty"`
}

func (r Response) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

//STOP1 OMIT

//START2 OMIT
func main() {
	response := Response{
		Message: "Hello",
	}

	// fmt.Printf wants Stringer's
	fmt.Printf("%v\n", response)
}

//STOP2 OMIT
