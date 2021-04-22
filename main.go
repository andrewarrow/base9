package main

import (
	"base9/nonary"
	"fmt"
)

func main() {
	thing := nonary.NewInt(0)
	for i := 0; i < 1000; i++ {
		thing = thing.Inc()
		fmt.Println(thing)
	}
}
