package main

import "fmt"

func main() {
	//string variable
	var name1 string = "golang"
	fmt.Println(name1)
	//infer
	var name2 = "golang"
	fmt.Println(name2)

	//boolean
	var isAdult = true
	fmt.Println(isAdult)

	//integer
	var age int = 30
	fmt.Println(age)

	//short hand syntax
	name3 := "Shankhalp"
	fmt.Println(name3)

	//if value is to be assigned later then the datatype needs to be entered
	var name4 string
	name4 = "Shankhalpah"
	fmt.Println(name4)

	var price float32 = 50.5
	fmt.Println(price)

	var price1 = 50.6
	fmt.Println(price1)

	price2 := 50.7
	fmt.Println(price2)

}
