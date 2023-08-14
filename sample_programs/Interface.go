
package main

import "fmt"

type Car interface{
	drive()
	stop()
}

type lambo struct{
	lambomodel string
}
type chevy struct{
	chevymodel string
}
func (l *lambo)drive(){
	fmt.Println("Lambo is in Drive")
	fmt.Println(l.lambomodel)
}

func anything(anything interface{}){

	fmt.Println(anything)
}

func (c *chevy) drive(){
	fmt.Println("Chevy is in Drive")
	fmt.Println(c.chevymodel)
}

func main(){
	l:= lambo{"Model L"}
	c:= chevy{"Model C"}
	l.drive()
	c.drive()
	anything("anjali")
	anything(610)
	mymap := make(map[string]interface{})
	mymap["Name"] = "Anjali"
	mymap["Age"] = 32
	fmt.Println(mymap)
}
