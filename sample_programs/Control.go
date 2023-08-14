package main

import "fmt"

func main() {

	fmt.Println("Controls in Go")
	f := "true"
	flag := &f

	if flag == nil {
		fmt.Println("Flag is Null")
	}
	else if *flag{
		fmt.Println("Flag is True")
	}
	else {
		fmt.Println("Flag is False")
	}

	arr := []string{"anjali", "katteri", "b22", "clyde", "court"}

	for i, value := range arr {
		fmt.Println(i)
		fmt.Println(value)
	}
	mymap := make(map[string]interface{})
	mymap["Anju"] = "28"
	mymap["anjali"] = 28

	for k, v := range mymap {
		fmt.Printf("name %s and age %v", k, v)
	}
}
