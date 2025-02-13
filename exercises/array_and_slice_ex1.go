package exercises

import "fmt"

func ArrayAndSliceEx1(){
	arr:= [5]int{1, 2, 3, 4 ,5}
	slice := arr[1:4]
	fmt.Printf("Array Original:\n %v \n",arr)
	fmt.Printf("Slice Original:\n %v \n", slice)
	arr[2] = 123
	slice[0] = 520
	fmt.Printf("Array Modified:\n %v \n",arr)
	fmt.Printf("Slice Modified:\n %v \n", slice)
}