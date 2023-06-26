package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	saltutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		saltutation = "welcome"
	}()

	wg.Wait()
	fmt.Println(saltutation)

}
