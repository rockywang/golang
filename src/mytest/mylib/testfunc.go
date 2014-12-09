package mylib

import (
	"fmt"
)

func TestMap() {
	rating := map[string]float32{"C": 5, "GO": 4.5, "Python": 4.5, "C++": 2}
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}
	delete(rating, "C")
	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m["Word"] = "leisong"
	m1 := m
	m1["Hello"] = "Salut"
	fmt.Println(m)
}
