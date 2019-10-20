package main

import "fmt"

func main() {
	//先执行后面的函数 在执行defer后面的函数
	defer fmt.Println("第一个函数")
	defer fmt.Println("第二个函数")
	defer fmt.Println("第三个函数")
	fmt.Println("第四个函数")
	fmt.Println("第五个函数")
	fmt.Println("第六个函数")
	fmt.Println("第七个函数")
	fmt.Println("第八个函数")
}

func ifelse() {
	b := false
	if b == true {
		fmt.Println("这是true")
	} else if b == false {
		fmt.Println("这是false")
	} else {
		fmt.Println("不知道是啥")
	}
}

func switchcase() {
	k := 20
	switch k {
	case 1:
		fmt.Println("这是1")
	case 2:
		fmt.Println("这是2")
	default:
		fmt.Println("什么都不是")
	}
}

func forfunc() {
	//for x < 10 {
	//	x++
	//	fmt.Println("打印机", x)
	//}
	//for x := 0; x < 10; x++ {
	//	fmt.Println("打印机", x)
	//}

	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	for i, e := range numbers {
		fmt.Println(e)
		fmt.Println(i)
	}

}
