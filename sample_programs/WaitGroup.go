package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		fmt.Println("First")
		wg.Done()
	}()
	go func() {
		fmt.Println("Second")
		wg.Done()
	}()
	fmt.Println("Before Main")
	wg.Wait()
	fmt.Println("Main")
}
