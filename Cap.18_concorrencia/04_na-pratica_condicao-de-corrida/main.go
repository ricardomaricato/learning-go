package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
- Aqui vamos replicar a race condition mencionada no artigo anterior.
    - time.Sleep(time.Second) vs. runtime.Gosched()
- go help → go help build → go run -race mainRob.go
- Como resolver? Mutex.
*/

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	contador := 0
	totaldegoroutines := 1000

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)

}
