package main

import "fmt"

func main() {
	var number int

	number = 10
	fmt.Println(number)

	var number1 int = 11
	fmt.Println(number1)
	//Type Inference: khi khai báo go sẽ tự hiểu biến bạn khai bao thuộc kiểu nào.

	var email = "kai@gmail.com"
	fmt.Println(email)

	//Khai báo nhiều biến
	//1. Khai báo nhiều biến với cùng kiểu dữ liệu
	var a, b int

	a = 1
	b = 2
	fmt.Println(a)
	fmt.Println(b)

	var a1, b1 int = 10, 11
	fmt.Println(a1)
	fmt.Println(b1)

	var a2, b2 = 12, 13
	fmt.Println(a2)
	fmt.Println(b2)

	//2. Khai báo nhiều biến khác kiểu dữ liệu
	var (
		name    string
		address string
		age     int
	)

	name = "Kaiz"
	address = "Viet Nam"
	age = 22

	fmt.Println(name)
	fmt.Println(address)
	fmt.Println(age)

	var (
		name1    string = "Mary"
		address1 string = "Germany"
		age1     int    = 21
	)
	fmt.Println(name1)
	fmt.Println(address1)
	fmt.Println(age1)

	var name2, address2, age2 = "Leon", "Ame", 23
	fmt.Println(name2)
	fmt.Println(address2)
	fmt.Println(age2)

	//Shorthand
	language := "Golang"
	fmt.Println(language)

}
