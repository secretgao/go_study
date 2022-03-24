package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T)  {
	/*
	单独声明
	var a int = 1
	var b int = 1
	一起声明
	var (
			a int = 1
			b  = 1
		)
	a:=1
	b:=1
	 */
	//一起声明
	 var a,b int = 1,1
	fmt.Print(a)
	for i:=0; i<5;i++{
		fmt.Print(" ",b)
		tmp:=a
		a=b
		b=tmp+a
	}
	fmt.Println()
}
