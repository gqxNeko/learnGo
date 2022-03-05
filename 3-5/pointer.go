package main

import "fmt"

func main() {
	var a int = 4
	var b int32 = 10
	var c float32 = 5.6
	var d uint = 5 
	fmt.Println(a,b,c,d)
	var ptr *uint
	ptr = &d
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
