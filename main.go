package main

import (
	"data-structure/exercises"
	"fmt"
)

func main() {
	exercises.WithGoRoutines()
	exercises.WithoutGoROutines()
	exercises.WithContext()

	fmt.Println("end of statement")
	
}



func printBoundCheck() {
	arr := [5]int{0,1,2,3,4}
	slice := arr[:]
	fmt.Println("Print Witho Hint Bound Check")
	// _ = slice[3] //bound checks
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])
	fmt.Println("-------------->")
	//run code with:
	//go run -gcflags="-d=ssa/check_bce"
}