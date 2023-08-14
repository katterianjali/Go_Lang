package main

import (
	"fmt"
)

type car struct {
	Name  string
	age   int64
	grade string
}

func (c car) Getname() string {
	return c.Name
}

func (c car) Details() {
	fmt.Println(c)
}

func main() {
	c := car{"Anjali", 28, "A"}
	c.Details()
	fmt.Println(c.Getname())
}
