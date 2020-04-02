package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"os"
)
//监控文件内容变化
func watch_log () {
	t, err := tail.TailFile("C:\\Users\\yyw\\go\\src\\main\\a.log", tail.Config{Follow: true})
	if err != nil {
		fmt.Println(err)
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}

func main()  {
	go watch_log()
	//打开文件
	f, err := os.OpenFile("C:\\Users\\yyw\\go\\src\\main\\a.log", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	//持续写入，for不断写入，起到使main程序不结束。go 程不会被退出
	for  {
		buf := []byte("test"+"\n")
		f.Write(buf)
	}
}

//大文件读写  bufio.NewReader  
