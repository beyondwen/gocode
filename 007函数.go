package main

import "fmt"

func main() {
	var a = add(1, 4)
	fmt.Println(a)
	var b, c = returnManyValue(1, "wenhao")
	fmt.Println(b, c)
	var d = variableParameter(1, 2, 3, 4)
	fmt.Println(d)
	var e, f = returnNamedValue()
	fmt.Println(e)
	fmt.Println(f)
	var g = recursive(0)
	fmt.Println(g)
	h := func() string {
		fmt.Println("函数值传递")
		return "函数调用"
	}
	fmt.Println(h)
	h()
	fmt.Println(transfer(h))
}

//返回单个值
func add(x int, y int) int {
	return x + y
}

//返回多个值
func returnManyValue(x int, y string) (int, string) {
	return x, y
}

//可变参数
func variableParameter(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//使用具名返回值
func returnNamedValue() (a, b string) {
	x := "100"
	y := "200"
	return x, y
}

//递归函数
func recursive(x int) int {
	total := x
	y := total + x
	y++
	if y < 1000 {
		return y
	}
	return recursive(y)
}

//将函数做为值传递
func transfer(f func() string) string {
	return f()
}
