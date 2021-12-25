package stack

import (
	"../generic-dynamic-array/array"
	"bytes"
	"fmt"
)

type ArrayStack struct {
	array *array.Array
}

func New(cap int)*ArrayStack{
	return &ArrayStack{
		array: array.New(cap),
	}

}

func (s *ArrayStack)GetSize()int{
	return s.array.GetSize()
}

func (s *ArrayStack)IsEmpty()bool{
	return s.array.IsEmpty()
}
//这个不是栈的特性，是用数组实现的栈的特性
func (s *ArrayStack)GetCapacity()int{
	return s.array.GetCap()
}

func (s *ArrayStack)Push(val interface{}){
	s.array.AddLast(val)
}

func (s *ArrayStack)Pop()interface{}{
	return s.array.RemoveLast()
}

//查看栈顶元素
func (s *ArrayStack)peek()interface{}{
	return s.array.GetLast()
}


//打印栈中被元素
func (s *ArrayStack)String()string{
	var buffer bytes.Buffer
	buffer.WriteString("stack:[")
	for i:=0;i<s.array.GetSize();i++{
		buffer.WriteString(fmt.Sprintf("%v",s.array.Get(i)))
		if i!=s.array.GetSize()-1{
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]top")
	return buffer.String()
}














