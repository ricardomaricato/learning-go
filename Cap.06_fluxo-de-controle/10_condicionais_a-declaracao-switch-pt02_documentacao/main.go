package main

import "fmt"

/*
- Pode avaliar tipos
- Pode haver uma expressão de inicialização
*/

var x interface{}

func main() {
	x = true

	switch x.(type) {
	case int:
		fmt.Println("é um int")
	case bool:
		fmt.Println("é um bool")
	case string:
		fmt.Println("é um string")
	case float64:
		fmt.Println("é um float")
	}
}
