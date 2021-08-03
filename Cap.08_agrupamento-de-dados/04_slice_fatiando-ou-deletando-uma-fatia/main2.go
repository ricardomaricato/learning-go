package main

import "fmt"

/*
- É fatiando que se deleta um item de uma slice. Na prática:
    - x := append(x[:i], x[:i]...)
*/

func main() {
	//		    0.           1.           2.         3.               4.
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marg"}

	sabores = append(sabores[:2], sabores[4:]...)

	fmt.Println(sabores)
}
