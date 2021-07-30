package main

import "fmt"

/*
- If, else.
- If, else if, else.
- If, else if, else if, ..., else.
*/

func main() {
	if x := 500; x > 100 {
		fmt.Println("chis é maior que cem")
	} else if x < 10 {
		fmt.Println("chis é menor que déis")
	} else {
		fmt.Println("chis não é menor que déis nem maior que cem")
	}
}
