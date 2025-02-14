package exercises

import (
	"fmt"
	"math"
)


func MapEx1() {
	 slice :=  make([]int, 3, 3);
	s := []int{1, 2, 3}
    s = append(s, 4, 5)
	slice2 := copy(slice, s)
	fmt.Println(slice2)
	dict := map[string]string{
		"Jhon" : "Doe",
		"Foo" : "Bar",
	}

	dict2 := map[string][]int{
		"Jhon" : {1, 2, 3},
		"Foo" : {9, 8 , 7},
	}

	fmt.Println(dict)
	fmt.Println(dict2)

	var dict3 = map[string]string{}
	dict3["Jhon"] = "Doe"
	value, ok := dict3["Jhon"]
	value2, ok2 := dict3["avbb"]

	fmt.Println(value, ok)
	fmt.Println(value2, ok2)
	
	delete(dict3, "Jhon")
	value, ok = dict3["Jhon"]
	fmt.Println(value, ok)

	clear(dict)
	fmt.Println(dict)

	dict4 := map[float64]string{}
	f1 := math.NaN()
	f2 := math.NaN()

	dict4[f1] = "Jhon"
	dict4[f2] = "Doe"
	fmt.Println(dict4)
	fmt.Println(f1 == f2)
	delete(dict4, f1)
	fmt.Println(dict4)
	value, ok = dict4[f2]
	fmt.Println(value, ok)
	clear(dict4)
	fmt.Println(dict4)

	dict = map[string]string{
		"Jhon" : "Doe",
		"Foo" : "Bar",
	}

	for key, value := range dict{
		fmt.Printf("key: %s\nvalue: %s \n", key, value)
		if key == "Jhon" {
			delete(dict, key)
		}
	}
	fmt.Println("--------------->")
	fmt.Println(dict)
}