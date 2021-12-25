/**
各种奇怪的实验
 */
package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	var index int=0
	var stackTest []int
	for i:=0;i<len(pushed);i++{
		stackTest=append(stackTest,pushed[i])
		for index<len(popped)&&len(stackTest)!=0&&stackTest[len(stackTest)-1]==popped[index]{
			stackTest=stackTest[:len(stackTest)-1]
			index++
		}
	}
	return len(stackTest)==0
}
func main(){
	var  arr []int=[]int{1,0}
	var arr2 []int=[]int{1,0}
	res:= validateStackSequences(arr,arr2)
	fmt.Println(res)

}

