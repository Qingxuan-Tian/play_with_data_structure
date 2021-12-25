package main

import (
	"../stacks-and-queues/stack"
	"fmt"
)
func main(){
	s:= stack.New(10)
	for i:=0;i<5;i++{
		s.Push(i)
		fmt.Println(s.String())
	}
	s.Pop()
	fmt.Println(s.String())
}
/**
output
stack:[0]top
stack:[0,1]top
stack:[0,1,2]top
stack:[0,1,2,3]top
stack:[0,1,2,3,4]top
stack:[0,1,2,3]top

 */
