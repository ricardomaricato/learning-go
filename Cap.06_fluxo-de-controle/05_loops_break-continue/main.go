package main

import "fmt"

/*
-continue quebra a iteração
-break para o loop
*/

func main() {
	for i := 0; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
