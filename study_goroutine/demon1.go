package main

import (
	"fmt"
	"sync"
)
var concurrenc int

func main()  {
	concurrenc = 10
	wg := sync.WaitGroup{}
	wg.Add(concurrenc+1)

	go findfilename(&wg)

	for i := 0;i < concurrenc; i++  {
		go func() {
			dosomework(&wg)
		}()
	}
	wg.Wait()
	fmt.Println("finish")
}

func findfilename(wg *sync.WaitGroup)  {
	wg.Done()
	fmt.Println("test")
}

func dosomework(wg *sync.WaitGroup)  {
	wg.Done()
	fmt.Println("do something")
}

// output:

// test
// do something
// do something
// do something
// do something
// do something
// do something
// do something
// do something
// do something
// finish
// do something

// Process finished with exit code 0
