package main

import "fmt"

func main() {
	x := 0

	for x < 10 {
		fmt.Println("chis é menor que déis")
		x++
	}

	y := 0

	for {
		if y < 10 {
			y++
			fmt.Println("chis é menor que déis")
		} else {
			fmt.Println("chis é maior que déis mano tô fora")
			break
		}
	}
}
