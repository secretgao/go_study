package type_test

import "testing"

type MyInt int64
func TestImplicit(t *testing.T )  {
	var  a int32 = 1
	var b int64
	var c MyInt
	//cannot use c (type MyInt) as type int64 in assignment 即使 相同类型 使用别名也不行
	//b = c
	t.Log(a,b,c)
}

func TestPoint(t *testing.T){
	a:=1
	aPrt:=&a
	//aPrt = aPrt+1  //不支持指针运算
	t.Log(a,aPrt)
	t.Logf("%T %T",a,aPrt)
}
func TestString(t *testing.T){
	var s string
	t.Log("*"+s+"*")
	t.Log(len(s))

	if s ==""{

	}
	/*  在go 判断空字符串 这种写法是错误的
	if s == nil{

	}

	 */
}