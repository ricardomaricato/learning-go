package main

import "fmt"

/*
- Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.
*/

func main() {
	c := make(chan int)
	go botalá(c)
	for v := range c {
		fmt.Println(v)
	}
}

func botalá(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
