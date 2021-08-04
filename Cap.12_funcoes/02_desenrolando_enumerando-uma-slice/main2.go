package main

import "fmt"

/*
- Pode-se passar *zero* ou mais valores
*/

func main() {

	//si := []int{10, 10, 1, 2, 3, 5}

	total := soma5()

	fmt.Println(total)
}

func soma5(x ...int) int {

	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}
