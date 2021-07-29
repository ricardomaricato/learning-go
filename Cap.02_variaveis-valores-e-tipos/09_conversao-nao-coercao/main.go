package main

import "fmt"

type hotdog int

var b hotdog = 20

func main() {
	x := 10
	fmt.Printf("%v\n", x)

	x = int(b)
	fmt.Printf("%v", b)
}
