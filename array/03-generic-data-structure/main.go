package main

import (
	"../03-generic-data-structure/array"
	"../03-generic-data-structure/student"
	"fmt"
)

func main(){
	//测试基本类型
	//hello world
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
	//测试student类型
	//根据名字找成绩，or根据成绩找名字，写额外的方法
	students := array.New(20)
	students.AddLast(student.New("xiaoming",99))
	students.AddLast(student.New("xiaohong",100))
	students.AddLast(1)
	fmt.Println(students.String())


}