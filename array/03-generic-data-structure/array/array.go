package array

import (
	"../../../utils/compare"
	"bytes"
	"fmt"
)

type Array struct {
	data []interface{}
	size int
}
//创建数组，有cap
func New(cap int)*Array{
	return &Array{
		data: make([]interface{}, cap),
		size: 0,
	}
}
//创建数组，没有定初始cap,默认为10
func NewWithNoCap()*Array{
	return &Array{
		data: make([]interface{}, 10),
		size:0,
	}
}
//获得容量大小
func (a *Array)GetCap()int{
	return len(a.data)
}

//获得已有元素个数
func (a *Array)GetSize()int{
	return a.size
}

//查看是否有元素
func (a *Array)IsEmpty()bool{
	return a.size==0
}
//指定位置插入元素
func (a *Array)Add(index int, e interface{}){
	if a.size==len(a.data){
		panic("array is full, can't add new element")
	}
	if index<0||index>a.size{
		panic("the index is wrong")
	}
	for i:=a.size-1;i>=index;i--{
		a.data[i+1]=a.data[i]
	}
	a.data[index]=e
	a.size++
}
/**在末尾插入元素,方法1
func(a *Array)AddLast(e int){
	if a.size==len(a.data){
		panic("数组已满")
	}
	a.data[a.size]=e
	a.size++
}

*/
//在末尾插入元素，方法2
func(a *Array)AddLast(e interface{}){
	a.Add(a.size,e)
}

//在数组头部添加元素
func(a *Array)AddFirst(e interface{}){
	a.Add(0,e)
}

/**获取数组中的特定元素
控制校验index的范围
通过这一步的封装，确保用户只能访问到已有的元素，不能访问到数组的其余部分
*/
func (a *Array)Get(index int)interface{}{
	if index<0|| index>=a.size{
		panic("get failed, index is illegal")
	}
	return a.data[index]
}
/**
修改index索引位置的元素为e
与get同理
*/
func (a *Array)Set(index int, e interface{}){
	if index<0||index>=a.size{
		panic("set failed, index is illegal")
	}
	a.data[index]=e
}

/**打印数组的详细信息
重写string的array方法
*/
func (a *Array)String() string{
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("array: size= %d, capacity= %d\n",a.size, len(a.data)))
	buffer.WriteString("[")
	for i:=0;i<a.size;i++{
		buffer.WriteString(fmt.Sprintf("%v", a.data[i]))
		if i!=(a.size-1){
			buffer.WriteString(",")
		}

	}
	buffer.WriteString("]")
	return buffer.String()
}
/**
查找数组中是否存在某个元素
*/
func (a *Array)Contains(e interface{})bool{
	for i:=0;i<a.size;i++{
		//if a.data[i]==e{
		//	return true
		//}
		if compare.CompareEqual(a.data[i],e){
			return true
		}
	}
	return false
}
/**
查找是否存在某个元素，如果存在，返回对应的第一个index，如果不存在，返回-1
如果有重复元素，只返回第一个
*/
func (a *Array)Find(e interface{})int{
	for i:=0;i<a.size;i++{
		//if a.data[i]==e{
		//	return i
		//}
		if compare.CompareEqual(a.data[i],e){
			return i
		}
	}
	return -1
}
/**
返回符合条件元素的index，以切片的形式
*/
func(a *Array)FindAll(e interface{})[]int{
	var res []int
	for i:=0;i<a.size;i++{
		//if a.data[i]==e{
		//	res=append(res,i )
		//}
		if compare.CompareEqual(a.data[i],e){
			res=append(res,i)
		}
	}
	return res
}

/**
从数组中删除对应index的元素
把要删除的节点的元素之后的元素向前覆盖
最后维护size
用户能看到的，需要关注的，也只是他向数组中添加过的元素，数组的其余部分，对用户不可见，用户也无需关心
*/
func (a *Array)Remove(index int)interface{}{
	if index<0||index>=a.size{
		panic("remove failed, index is illegal")
	}
	ret:= a.data[index]
	for i:=index+1;i<a.size;i++{
		a.data[i-1]=a.data[i]
	}
	a.size--
	return ret

}
/**
删除第一个元素
*/
func (a *Array)RemoveFirst()interface{}{
	return a.Remove(0)
}
/**
删除最后一个元素
*/
func (a *Array)RemoveLast()interface{}{
	return a.Remove(a.size-1)
}

/**
从数组中删除对应的元素
若有重复元素，只删除第一个元素
*/
func (a *Array)RemoveElement(e interface{})bool{
	index:= a.Find(e)
	if index!=-1{
		a.Remove(index)
		return true
	}
	return false
}
/**
删除全部对应的元素
*/
func (a *Array)RemoveAllElement(e interface{})bool{
	indexs:= a.FindAll(e)
	if len(indexs)==0{
		return false
	}
	for i:=0;i<a.size;i++{
		//if a.data[i]==e{
		//	a.Remove(i)
		//}
		if compare.CompareEqual(a.data[i],e){
			a.Remove(i)
		}
	}
	return true
}



