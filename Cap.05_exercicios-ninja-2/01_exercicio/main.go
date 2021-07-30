package main

import "fmt"

/*
- Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
*/

func main() {
	x := 100
	fmt.Println("binary\t\tdecimal\t\thexadecimal")
	fmt.Printf("%d\t\t", x)
	fmt.Printf("%#x\t\t", x)
	fmt.Printf("%b\n", x)
}
