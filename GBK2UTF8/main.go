package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"sync"
)


var rar_path string = "/data/zip/unrarfull"
var filename = make(chan string )

func main() {
	n := 30
	wg := sync.WaitGroup{}
	wg.Add(1)
	go GetAllFile(rar_path,filename)
	for i := 0; i < n ; i++  {
		go func() {
			for {
				filenames := <- filename
				fmt.Println(filenames)
				dpath := path.Dir(filenames) + "/done."+path.Base(filenames)
				unrar(filenames,dpath)
			}
		}()
	}
	wg.Wait()
	fmt.Println("finished")
}

func unrar(fn string,dpath string)  {
	cmd := exec.Command("/usr/bin/iconv","-f" ,"gbk" ,"-t" ,"utf-8",fn,"-o",dpath)
	buf, _ := cmd.Output()          // 错误处理略
	print(string(buf))
	remove_file(fn)
}

func GetAllFile(pathname string,fn_ch chan string) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			GetAllFile(fullDir,fn_ch)
			if err != nil {
				fmt.Println("read dir fail:", err)
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			fn_ch <- fullName
		}
	}
}

func remove_file(file string)  {
	err := os.Remove(file)               //删除文件test.txt
	if err != nil {
		//如果删除失败则输出 file remove Error!
		fmt.Println("file remove Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	} else {
		//如果删除成功则输出 file remove OK!
		fmt.Print("file remove OK!")
	}
}
