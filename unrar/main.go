package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"sync"
)

var unrar_path string = "/data/zip/unrarfull"
var rar_path string = "/data/zip/full"
var filename = make(chan string )
var donech =make(chan int)

func main() {
	n := 30
	wg := sync.WaitGroup{}
	wg.Add(1)
	go GetAllFile(rar_path,filename)
	for i := 0; i < n ; i++  {
		go func() {
			for {
				filenames := <- filename
				dpath := unrar_path + "/" + path.Base(filenames)
				mkdir(dpath)
				unrar(filenames,dpath)
			}
		}()
	}
	wg.Wait()
	fmt.Println("finished")
	//time.Sleep(10*time.Second)
}

func unrar(fn string,dpath string)  {
	cmd := exec.Command("/usr/local/bin/unrar","e",fn,dpath)
	buf, _ := cmd.Output()          // 错误处理略
	print(string(buf))
}

func GetAllFile(pathname string,fn_ch chan string) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
	}

	fmt.Println(rd)
	for _, fi := range rd {
		if fi.Size() == 391 {
			fmt.Println(fi.Name()+"is not RAR archive")
		}else {
			fullName := pathname + "/" + fi.Name()
			fn_ch <- fullName
		}
	}
}

func mkdir(n string)  {
	err := os.MkdirAll(n, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}
