package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println(os.Args)
	if len(os.Args) > 1{
		fmt.Println("hello",os.Args[1])
	} else {
		os.Exit(-1)
	}
}
