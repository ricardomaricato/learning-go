package main

import "fmt"

func main() {
	x := "oi bom dia\ncomo vai?\tespero \"que\" esteja tudo bem."
	fmt.Println(x)

	y := "oi"
	z := "bom dia"

	a := fmt.Sprint(y, " ", z)
	fmt.Println(a)
}
