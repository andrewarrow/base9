package main

import (
	"base9/nonary"
	"fmt"
)

func main() {
	thing := nonary.NewInt(0)
	for i := 0; i < 100; i++ {
		thing = thing.Add(1)
		fmt.Println(thing)
	}
}
