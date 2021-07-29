package main

import "fmt"

func main() {
	x := 10.0
	y := "olá bom dia"
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)

	a := 10 < 20
	fmt.Printf("a: %v, %T\n", a, a)
}
