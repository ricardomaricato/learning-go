package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(2)
	go func() { fmt.Println("Eu sou a goroutine número: 1"); waitGroup.Done() }()
	go func() { fmt.Println("Eu sou a goroutine número: 2"); waitGroup.Done() }()
	waitGroup.Wait()
}
