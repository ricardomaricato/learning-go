package main

import "fmt"

/*
- Com loops.
*/

func main() {
	fmt.Println(fatorial2(6))
	fmt.Println(loops(6))

}

func fatorial2(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}

func loops(x int) int {
	total := 1
	for x > 1 {
		total *= x // total = total * x
		x--
	}
	return total
}
