package array

type Array struct {
	data []int
	size int
}
//创建数组，有cap
func New(cap int)*Array{
	return &Array{
		data: make([]int, cap),
		size: 0,
	}
}
//创建数组，没有定初始cap,默认为10
func NewWithNoCap()*Array{
	return &Array{
		data: make([]int, 10),
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
