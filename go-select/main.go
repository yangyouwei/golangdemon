package main

import (
	"fmt"
	"time"
)

//用法2 超时控制  time.after
func main() {
	ch := make(chan int)

	select {
	case <- ch :
		fmt.Println("ch")
	case <- time.After(1):
		fmt.Println("timeout")
	}
}

//用法1
//func main() {
//	ch := make(chan int)
//	go func() {
//		ch <- 1
//	}()
//
//	time.Sleep(time.Second) //如果存在default必须有睡眠，没有睡眠的话，goroutine没运行就执行到了default分支
//
//	select {
//	case <- ch :
//		fmt.Println("ch")
//	default:
//		fmt.Println("default")
//	}
//	fmt.Println("test")
//}

//用法2  timeout
//func main() {
//	ch := make(chan int)
//	timeout := make(chan int,1)
//	go func() {
//		time.Sleep(time.Second)
//		timeout <- 1
//	}()
//
//	select {
//	case <- ch :
//		fmt.Println("ch")
//	case <- timeout:
//		fmt.Println("timeout")
//	}
//}


//程序错误  fatal error: all goroutines are asleep - deadlock!
//func main() {
//	ch := make(chan int)
//	select {
//	case <- ch :   //一直被阻塞,导致主routine也被阻塞，一般都有个default
//		fmt.Println("ch")
//	}
//	fmt.Println("test")
//}


//select只循环一次
//一般用select放到for循环中，监控事件。
//func ()  {
//	for  {
//		select {
//		case :
//			dosmothing
//		case :
//			dosmothing
//		case :
//			dosmothing
//		default:
//			dosmothing
//		}
//	}
//}

