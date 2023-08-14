package main

import (
	"fmt"
	"time"
)

type car struct {
	Model string
}

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	b := make(chan int, 3)
	go func() {
		b <- 1
		b <- 2
		b <- 3
		b <- 4
		close(b)
	}()
	m := make(chan *car, 3)
	go func() {
		m <- &car{"M1"}
		m <- &car{"M2"}
		m <- &car{"M3"}
		m <- &car{"M4"}
		close(m)
	}()
	val := <-c
	fmt.Println(val)

	for i := range b {
		fmt.Println(i)
	}
	for j := range m {
		fmt.Println(j.Model)
	}
	time.Sleep(time.Second * 2)
	fmt.Println(c)

}
