package ch05

// learn pointers
func LearnPointers() {
	// declare a variable
	var a int = 10

	// declare a pointer
	var p *int

	// assign the address of a to p
	p = &a

	// print the value of a
	println(a)

	// print the address of a
	println(&a)

	// print the value of p
	println(p)

	// print the value of a through p
	println(*p)

	// change the value of a through p
	*p = 20

	// print the value of a
	println(a)
}
