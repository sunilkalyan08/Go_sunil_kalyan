package main

import (
	"fmt"
	
)

func main() {
	fmt.Println("Welcome! What's your name?")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println("why ur here?")
	var ques string
	fmt.Scanf("%s", &ques)
	fmt.Println("Welcome! What's your age?")
	var age int
	fmt.Scanf("%d", &age)
	// fmt.Printf("Details Entered:")
	// fmt.Printf("Name,",name)
	// fmt.Printf("Age,",age)
	// fmt.Printf("Reason,",ques)
	fmt.Printf("Nice to have you onboard, %s (%d) is for %s", name,age,ques)

}