package main

import "fmt"

type movie struct {
	name string
	age  int
}

type jack struct {
	name name
}

type name struct {
	a int
	b int
	c int
}

func main() {
	m := movie{
		name: "wenhao",
		age:  26,
	}
	fmt.Println(m.age)
	fmt.Println(m.name)

	var x movie
	x.name = "test"
	x.age = 12
	fmt.Println(x)

	y := new(movie)
	y.name = "hhahhah"
	y.age = 12
	fmt.Println(y)

	g := jack{
		name: name{a: 1, b: 2, c: 2},
	}
	fmt.Println(g.name.a)

	r := name{a: 1, b: 2, c: 3}
	j := name{a: 2, b: 2, c: 3}
	if r == j {
		fmt.Println("比较结构")
	}

}

func newJack(x int) jack {
	a := jack{name: name{a: 1, b: 2, c: 3}}
	return a
}
