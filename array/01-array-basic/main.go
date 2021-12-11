package main

import "fmt"

func main(){
	var arr [10]int
	for i:=0;i<len(arr);i++{
		arr[i]=i
	}
	arr2:= [...]int{99,98,66}
	for i:=0;i<len(arr2);i++ {
		fmt.Println(arr2[i])
	}
	for _, val:= range arr2{
		fmt.Println(val)
	}
	arr2[0]=100
	fmt.Println(arr2)
}
