package ch02

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func PrintDataTypes() {

	var age int
	age = 29
	fmt.Println("Age: ", age)

	var name string
	name = "Nitin"
	fmt.Println("Name: ", name)

	var isCool bool
	isCool = true
	fmt.Println("isCool: ", isCool)

	var size float32
	size = 2.3
	fmt.Println("size: ", size)

	// Shorthand
	age2 := 30
	fmt.Println("Age2: ", age2)

	name2 := "Nitin"
	fmt.Println("Name2: ", name2)

	isCool2 := true
	fmt.Println("isCool2: ", isCool2)

	// Slices
	var emails []string
	emails = []string{"nitin@gmail.com", "nitin2@gmail.com"}
	fmt.Println("Emails: ", emails)

	var hobbies = [2]string{"reading", "coding"}
	fmt.Println("Hobbies: ", hobbies)

	var amounts = []float32{2.3, 4.5}
	fmt.Println("Amounts: ", amounts)

	// Maps (key-value pairs)
	var address = map[string]string{"city": "Bangalore", "country": "India"}
	fmt.Println("Address: ", address)

	// Any type
	var a interface{} = 23
	fmt.Println("a: ", a)

	var b interface{} = "Nitin"
	fmt.Println("b: ", b)

}

func PrintPersonStruct() {
	person1 := Person{FirstName: "Nitin", LastName: "Sharma", Age: 29}
	fmt.Println("Person1: ", person1)
	fmt.Println("Person1 FirstName: ", person1.FirstName)
	fmt.Println("Person1 LastName: ", person1.LastName)
	fmt.Println("Person1 Age: ", person1.Age)

	var person2 Person
	person2.FirstName = "Nitin"
	person2.LastName = "Sharma"
	person2.Age = 29
}
