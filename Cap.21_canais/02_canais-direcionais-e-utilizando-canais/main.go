package main

import "fmt"

/*
Canais direcionais

- Canais podem ser direcionais.
- E isso serve pra...?
- Um send channel e um receive channel são tipos diferentes. Isso permite que os type-checking mechanisms do compilador façam com que não seja possível, por exemplo, escrever num canal de leitura.
- Aos aventureitos: https://stackoverflow.com/questions/1...
- Canais bidirecionals (send & receive)
    - send chan←
        - error: "invalid operation: ←cs (receive from send-only type chan← int)"
    - receive ←chan
        - error: "invalid operation: cr ← 42 (send to receive-only type ←chan int)"
*/

func main() {
	canal := make(chan int)

	go send(canal)

	receive(canal)

}

func send(s chan<- int) {
	s <- 42
}

func receive(r <-chan int) {
	fmt.Println("O valor recebido do canal foi:", <-r)
}
