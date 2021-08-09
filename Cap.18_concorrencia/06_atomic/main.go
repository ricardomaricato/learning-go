package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
- Agora vamos fazer a mesma coisa, mas com atomic ao invés de mutex.
    - atomic.AddInt64
    - atomic.LoadInt64
*/

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())

	var contador int64
	contador = 0
	totaldegoroutines := 15

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Valor final:", contador)

}
