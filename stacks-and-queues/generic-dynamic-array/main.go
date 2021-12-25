package main

import (
	".//array"
	"fmt"
)

func main(){
	//测试基本类型
	//hello world
	a:=array.New(5)
	for i:=0;i<5;i++{
		a.AddLast(i)
	}
	fmt.Println(a.String())
	a.Add(1,10)
	fmt.Println(a.String())
	a.RemoveAllElement(10)
	a.RemoveLast()
	a.RemoveLast()
	fmt.Println(a.String())
	//测试student类型
	//根据名字找成绩，or根据成绩找名字，写额外的方法
	//students := array.New(20)
	//students.AddLast(student.New("xiaoming",99))
	//students.AddLast(student.New("xiaohong",100))
	//students.AddLast(1)
	//fmt.Println(students.String())


}