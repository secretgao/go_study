package constant_test

import "testing"

//连续常量简洁写法
const  (
	Monday = iota + 1
	Tuesday
	Wednesday
)
//
const  (
	Readable = 1 <<iota
	Writable
	Executable
)
func TestConstant(t *testing.T){
	t.Log(Wednesday,Tuesday,Monday)
	t.Log(Readable,Writable,Executable)
}
func TestConstant1(t *testing.T){
	a:=7

	t.Log(a&Readable==Readable,a&Writable==Writable,a&Executable==Executable)
}