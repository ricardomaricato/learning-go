package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	funções := 5
	go manda2(100, canal1)
	go outra2(funções, canal1, canal2)
	for v := range canal2 {
		fmt.Println(v)
	}
}

func manda2(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func outra2(funções int, canal1, canal2 chan int) {
	var wg sync.WaitGroup
	for i := 0; i < funções; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho2(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)
}

func trabalho2(n int) int {
	time.Sleep(time.Millisecond * 1000) //time.Duration(rand.Intn(1e3)))
	return n
}
