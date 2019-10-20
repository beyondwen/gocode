package main

import "fmt"

func main() {
	//var 形式声明变量
	var s string
	s = "hello word \n"
	fmt.Print(s)

	//声明多个变量并赋值
	var x, y string = "q\n", "b\n"
	fmt.Print(x)
	fmt.Print(y)

	//未赋值的变量 默认值
	var z string
	var j bool
	var k int64
	fmt.Print(j)
	fmt.Print(z)
	fmt.Print(k)
	if "" == z {
		fmt.Print("我是默认\n")
	}
	fmt.Print("===========\n")

	//简短变量声明
	q := 100
	fmt.Print(q)

}
