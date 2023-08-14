package main

import (
	"fmt"
	"time"
)

func waiting() {
	time.Sleep(time.Second * 5)
}
func wait1() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("waited 1 sec")
	}
}

func wait3() {
	for {
		time.Sleep(time.Second * 3)
		fmt.Println("waited 3 seconds")
	}
}

func main() {
	waiting()
	go waiting()
	go wait1()
	go wait3()
	fmt.Println("finished")
	select {}
}
