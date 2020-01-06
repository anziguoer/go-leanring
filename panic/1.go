package main

import "fmt"

func recoverName()  {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}

func fullName(firstName *string, lastName *string)  {
	//defer fmt.Println("deferred call in fullName ====== 1")
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}

	if lastName == nil {
		panic("runtime error: last name cannot be nil ===== 3")
	}

	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullname")
}

func main() {
	defer fmt.Println("deferred call in main === 2")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
