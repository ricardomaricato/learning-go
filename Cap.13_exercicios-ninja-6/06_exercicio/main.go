package main

import (
	"fmt"
)

/*
- Crie e utilize uma função anônima.
*/

func main() {

	func() {
		fmt.Println("Hello, playground")
	}()
}
