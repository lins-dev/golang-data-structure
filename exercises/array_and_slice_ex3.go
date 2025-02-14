package exercises

import "fmt"


func ArrayAndSliceEx3() {
	arr := [5]int{0,1,2,3,4}
	slice := arr[:]
	tryChangeArray(arr)
	fmt.Println(arr)

	tryChangeSlice(slice)
	fmt.Println(slice)
	
	tryChangeArrayPointer(&arr)
	fmt.Println(arr)

}


func tryChangeArray(arr [5]int) {
	arr[1] = 55
	
}

func tryChangeSlice(slice []int) {
	slice[1] = 42
	
}

func tryChangeArrayPointer(arr *[5]int) {
	arr[2] = 51
	
}