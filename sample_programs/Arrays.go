package main

import (
	"fmt"
	"strings"
)

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []string{"hello", "world"}
	fmt.Println(arr1, arr2)
	arr1 = append(arr1, 5)
	fmt.Println(arr1, arr2)
	fmt.Println(arr1, strings.Split(arr2[0], "e"))
}
