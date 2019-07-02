package main

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)


//目录操作取父目录
//取basename
func main()  {
		path_name := "./asdf/asdf/asd/f3fa/sdggsdf/a.log"
		e,_ := filepath.Abs(path_name)  //转换绝对路径
		a := strings.LastIndex(e, "\\")  //如果是linux系统使用 / 为分隔符
		rs := []rune(e)
		pname := rs[:a]  //父目录
		fmt.Println(pname)
		fmt.Println(path.Base(path_name)) //basename

}



//移动文件
//func main() {
//	oldPath, newPath := "C:\\a.log", "C:\\test\\a.log"  //必须是文件全路径带文件名的目标目录不能只是个目录名，不改名需要带上原文件名，并且移动支持同一个磁盘上移动。
//	err := os.Rename(oldPath, newPath)
//	if err != nil {
//		log.Panic(err)
//	}
//
//	//删除目录
//	err = os.RemoveAll("D:\\test")
//	if err != nil {
//		log.Panic(err)
//	}
//}


//创建多级目录
//func main()  {
//	//创建多级目录
//	err :=os.MkdirAll("./dir1/dir2",os.ModePerm)
//	if err!=nil{
//		fmt.Println(err)
//	}
//}

//文件校验md5
//func main()  {
//	filepaht := "D:\\迅雷下载\\disk-1.vmdk"   //文件大小 8.32G
//	bbb(&filepaht)
//}



////文件校验 使用io.copy效率高点。但是随着文件大小变化，校验时间响应变化。
//func bbb(path *string) {
//	f, err := os.Open(*path)
//	filename := filepath.Base(*path)
//	if err != nil {
//		fmt.Println("Open", err)
//		return
//	}
//
//	defer f.Close()
//
//	md5hash := md5.New()
//	t := time.Now()
//	if _, err := io.Copy(md5hash, f); err != nil {
//		fmt.Println("Copy", err)
//		return
//	}
//	elapsed := time.Since(t)
//	md5hash.Sum(nil)
//	fmt.Printf("time : %v ,filename: %s , md5: %x\n",elapsed,filename,md5hash.Sum(nil))
//}
