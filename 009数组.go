package main

import "fmt"

func main() {

	//数组
	var array [10]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array)
	//切片
	var chess = make([]string, 3)
	chess[0] = "1"
	chess[1] = "4"
	chess = append(chess, "wertyu")
	chess = append(chess, "hhhhhh", "dddddd", "eeeeeee")
	//chess = append(chess[:3], chess[4:]...)
	fmt.Println(chess)
	fmt.Println(len(chess))
	var mild = make([]string, 2)
	copy(mild, chess)
	fmt.Println(mild)
}
