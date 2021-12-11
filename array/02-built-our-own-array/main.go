package main

import (

	"../02-built-our-own-array/array"
	"fmt"

)
func main(){
	a:=array.New(5)
	fmt.Println(a)
	fmt.Println(a.IsEmpty(),a.GetCap(),a.GetSize())
}
