package main

import "fmt"

func main() {
	s := 100
	//传递指针的方式
	showAddr(&s)
	fmt.Println(&s)
}

//传递指针的方式
func showAddr(s *int) {
	fmt.Println(s)
}
