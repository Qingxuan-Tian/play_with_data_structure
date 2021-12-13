package main

import (
	"../02-built-our-own-array/array"
	"fmt"
)

func main(){
	a:=array.New(20)
	for i:=0;i<10;i++{
		a.AddLast(i)
	}
	a.Add(1,10)
	a.Set(2,99)
	a.AddLast(99)
	//[0,10,99,2,3,4,5,6,7,8,9,99]
	//fmt.Println(a.String())
	fmt.Println(a.Find(99))
	fmt.Println(a.FindAll(99))
	a.RemoveAllElement(99)
	fmt.Println(a.String())


}
