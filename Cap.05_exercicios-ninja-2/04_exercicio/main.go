package main

import "fmt"

/*
- Crie um programa que:
- Atribua um valor int a uma variável
- Demonstre este valor em decimal, binário e hexadecimal
- Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
- Demonstre esta outra variável em decimal, binário e hexadecimal
*/

func main() {
	x := 10
	fmt.Println("binary\t\tdecimal\t\thexadecimal")
	fmt.Printf("%d\t\t", x)
	fmt.Printf("%#x\t\t", x)
	fmt.Printf("%b\n", x)

	y := x << 1
	fmt.Println("binary\t\tdecimal\t\thexadecimal")
	fmt.Printf("%d\t\t", y)
	fmt.Printf("%#x\t\t", y)
	fmt.Printf("%b\n", y)
}
