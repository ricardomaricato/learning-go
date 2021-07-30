package main

import "fmt"

/*
- If: bool
- If: o operador não → "!"
- If: declaração de inicialização
*/

func main() {
	x := 10
	if !(x > 100) {
		fmt.Println("Hello, playground")
	}
}
