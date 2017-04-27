package main 

import (
	"bublesort/algo"
	"fmt"
)

func main() {
	array:=[]int{12, 13, 50, 17, 111, 103}
	PrintArray(array)
	algo.BubleSort(array)
	PrintArray(array)
}

func PrintArray(array []int) {
	fmt.Println()
	for num:= range array {
		fmt.Print(array[num])
		fmt.Print(" ")
	} 
}