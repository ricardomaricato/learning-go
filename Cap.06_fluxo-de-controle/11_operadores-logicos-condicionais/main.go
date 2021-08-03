package main

import "fmt"

/*
- &&
- ||
- !
- Qual o resultado de fmt.Println...
    - true && true
    - true && false
    - true || true
    - true || false
    - !true
*/

func main() {
	x := 9

	if !(x%2 == 0) && x%3 == 0 {
		fmt.Println("é múltiplo de dois e tambem de treis")
	}

	//	if x%2 == 0 || x%3 == 0 {
	//		fmt.Println("é múltiplo de dois ou de treis")
	//	}
}
