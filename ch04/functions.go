package ch04

import "errors"

type MyError error // custom error type

func LearnFunctions() {
	// function with no return value
	sayHello()

	// function with return value
	sum := add(1, 2)
	println(sum)

	// function with multiple return values
	result, err := divide(10, 2)
	if err != nil {
		println(err.Error())
		return
	}
	println(int(result))

	// function with named return values
	division, err := divideNamed(10, 0)
	if err != nil {
		println(err.Error())
		return
	}
	println(division)
}

func sayHello() {
	println("Hello")
}

func add(a, b int) int {
	return a + b
}

func divide(a float32, b float32) (float32, error) {
	if b == 0 {
		myError := MyError(errors.New("cannot divide by zero"))
		return 0, myError
	}
	return a / b, nil
}

func divideNamed(a int, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("cannot divide by zero")
		return 0, err
	}
	result = a / b
	return result, nil
}
