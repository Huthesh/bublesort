package algo

import (

)


func BubleSort(array []int){
	
	length:=len(array)
	for i:=0;i<length;i++{
		for j:=0;j<length-i-1;j++{
			if(array[j+1]>array[j]){
					temp:=array[j]
					array[j]=array[j+1]
					array[j+1]=temp
				}
		}
	}
}