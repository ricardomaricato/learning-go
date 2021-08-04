package main

import "fmt"

/*
- Mais um:
*/

func main() {

	total, quantos, oi := soma3("tarde", 10, 10, 1, 2, 3, 5)

	fmt.Println(total, quantos, oi)
}

func soma3(s string, x ...int) (int, int, string) {
	oi := ""
	if s == "manhã" {
		oi = "Oi, bom dia!"
	} else if s == "tarde" {
		oi = "Oi, boa tarde!"
	} else {
		oi = "Oi, boa noite!"
	}
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), oi
}
