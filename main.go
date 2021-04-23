package main

import (
	"base9/nonary"
	"fmt"
)

func main() {
	thing := nonary.NewInt()
	for i := 0; i < 100; i++ {
		thing.Add(1)
		fmt.Println(thing.String())
	}
}
