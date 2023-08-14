package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	day := "Monday"
	switch day {
	case "friday":
		fmt.Println("Its weekend")
	case "Monday", "tuesday":
		fmt.Println("Its working day")
	default:
		fmt.Println("Nearing weekend")
	}
}
