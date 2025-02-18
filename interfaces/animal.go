package interfaces

import "fmt"



type Animal interface{
	MakeSound() string
}


func TakeAnimal (a Animal) string {
	switch t := a.(type) {
	case *Dog:
		fmt.Println("This a Dog")
		return t.MakeSound()
	case *Cat:
		fmt.Println("This a Cat")
		return t.MakeSound()
	default:
		return "Animal unknow!"
	}
}