package ch03

func LearnControlStructures() {
	// if else
	if 1 > 0 {
		println("1 is greater than 0")
	} else {
		println("1 is not greater than 0")
	}

	// switch
	number := 1
	switch number {
	case 0:
		println("0")
	case 1:
		println("1")
	default:
		println("default")
	}

	// for loop
	for i := 0; i < 5; i++ {
		println(i)
	}

	// range
	nums := []int{11, 12, 13}
	for i, num := range nums {
		println(i, num)
	}

	// break
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		println(i)
	}

	// continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		println(i)
	}

}
