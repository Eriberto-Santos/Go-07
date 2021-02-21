package main

import (
	"./greet"
	"fmt"
)

func main() {
	result := greet.English()
	result2 := greet.Spanish()
	fmt.Println(result)
	fmt.Println(result2)
}
