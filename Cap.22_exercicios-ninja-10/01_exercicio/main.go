package main

import "fmt"

/*
- Nível 10?! Êita! Parabéns!
- Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
    - Usando uma função anônima auto-executável
    - Usando buffer
*/

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
