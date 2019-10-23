package main

import "fmt"

type wenhao struct {
	name string
	age  int
}

func main() {
	a := wenhao{name: "wenhao", age: 12}
	fmt.Println(a)
	b := a
	fmt.Println(b)
	b.name = "lisi"
	b.age = 12
	fmt.Println(b)
	c := &a
	fmt.Println(c)
	c.name = "zhangsan"
	c.age = 16
	fmt.Println(c)
	fmt.Println(a)
}
