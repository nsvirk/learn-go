package ch02

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

func MarshalJson() {

	// populate the struct Animal
	dog := Animal{Name: "Dog", Age: 3, Color: "Brown"}
	fmt.Println(dog)

	// marshal Go struct to json
	jsonDog, err := json.Marshal(dog)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonDog))

	// unmarshal json to Go struct
	var cat Animal
	jsonCat := []byte(`{"name":"Cat","age":2,"color":"White"}`)
	err = json.Unmarshal(jsonCat, &cat)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cat)

}
