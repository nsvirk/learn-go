package ch01

import "fmt"

func LearnFmt() {
	fmt.Println("from LearnFmt")
}

func LearnFmtPrintf() {
	fmt.Printf("printing from %s which has %v characters \n", "LearnFmtPrintf", len("LearnFmtPrintf"))
}

func LearnFmtSprintf() {
	var s string = ""
	s = fmt.Sprintf("printing from %s which has %v characters \n", "LearnFmtSprintf", len("LearnFmtSprintf"))
	fmt.Println(s)
}
