package main

import (
	"fmt"
	"github.com/Eriberto-Santos/Go-07/greet"
)

func main() {
	result := greet.English()
	result2 := greet.Spanish()
	fmt.Println(result)
	fmt.Println(result2)
}
