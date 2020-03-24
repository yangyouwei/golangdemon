package main

import (
	"fmt"
	"reflect"
)

//make 只能对slice map  chanle 操作。其他的用new

func main() {
	mp := make(map[string]string,5)
	fmt.Printf("type:%v\n",reflect.TypeOf(mp))
	// 此处输出map[string]string
	sl := make([]string, 2, 10)
	fmt.Printf("type:%v\n",reflect.TypeOf(sl))
	// 此处输出[]string
	sn := new(int)
	fmt.Printf("ntp : %v\n", reflect.TypeOf(sn))
	// 此处输出*int
}
