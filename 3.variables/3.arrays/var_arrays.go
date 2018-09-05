package main

import "fmt"

func main() {
	// Arrays length is a part of its type. Arrays cannot be resized.
	var vals [3]int
	vals[0] = 2
	vals[1] = 4
	vals[2] = 6

	fmt.Println(vals)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)

}
