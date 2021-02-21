package main

import (
	"fmt"

	"github.com/Eriberto-Santos/Go-07/greet"
	"rsc.io/quote"
)

func main() {
	result := greet.English()
	result2 := greet.Spanish()
	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(quote.Hello())
}
