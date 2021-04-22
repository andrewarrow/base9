package main

import (
	"base9/nonary"
	"fmt"
)

func main() {
	thing := nonary.NewInt(0)
	thing.Add(101)
	thing.Add(-91)
	fmt.Println(thing)
}
