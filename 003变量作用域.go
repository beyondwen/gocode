package main

import "fmt"

func main() {
	//内部块可以访问外部块
	//外部快无法访问内部块
	a := 100
	fmt.Println(a)
	{
		b := 1000
		fmt.Println(b)
		fmt.Println(a)
		{
			c := 10000
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
		}
	}
}
