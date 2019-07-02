package main

import (
	"fmt"
)

func watch(ch chan int) {
	for {
		ch <- 999
		fmt.Println("test")//不知道为啥必须干点啥。否则内存疯长，一般用上通道了都会用通道干点啥的。
		//time.Sleep(time.Duration(3)*time.Second)
	}
}

func dowork(a int) {
	fmt.Println(a)
	return
}

func main() {
	var ch = make(chan int, 100)
	go watch(ch)
	//for循环使main不退出
	for {
		a := <-ch //通过chanle 阻塞for循环否则无限启动goroutine
		go dowork(a)
	}
}
