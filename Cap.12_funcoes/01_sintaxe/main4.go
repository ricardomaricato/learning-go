package main

import "fmt"

/*
- Função com múltiplos retornos e parâmetro variádico.
*/

func main() {

	total, quantos, oi := soma2(10, 10, 1, 2, 3, 5)

	fmt.Println(total, quantos, oi)
}

func soma2(x ...int) (int, int, string) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), "Bom dia!"
}
