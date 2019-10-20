package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	s := "true"
	var b bool
	b, _ = strconv.ParseBool(s)
	fmt.Print(reflect.TypeOf(b))
}
