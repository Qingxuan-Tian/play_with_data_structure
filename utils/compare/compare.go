package compare

import "reflect"

/**
比较两个接口是否相等
 */
func CompareEqual(a, b interface{})bool{
	//判断类型是否一样，不一样抛出异常
	if reflect.TypeOf(a).Kind()!= reflect.TypeOf(b).Kind(){
		panic("compare failed, the type of two interface is different")
	}
	//默认传进来的类型都是可比较的，不可比较的会触发panic
	return a==b
}
