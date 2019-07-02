package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func spiderpage(i int,p chan int )  {
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn="+strconv.Itoa((i-1)*50)
	fmt.Println("正在爬取第 %d 个页面",i)
	result, err := HttpGET(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	filename := strconv.Itoa((i-1)*50)+".html"
	f, err  := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.WriteString(result)
	f.Close()
	p <- i
}

func Dowrk(start, end int)  {
	fmt.Printf("正在爬取 %d 到 %d 的页面\n",start,end)
	page := make(chan int)
	for i := start; i <= end  ; i++ {
		go spiderpage(i,page)
	}
	for i := start; i <= end  ; i++ {
		fmt.Println("第 %d 个页面爬取完成",<-page)
	}
}

func HttpGET(url string)(result string,err error)  {
	reps, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer reps.Body.Close()
	buf := make([]byte,1024*4)
	for   {
		n, err := reps.Body.Read(buf)
		if n == 0 {
			fmt.Println(err)
			break
		}
		result += string(buf[:n])
	}
	return
}

func main(){
	var star, end int
	fmt.Printf("请输入起始页 （>=1） :")
	fmt.Scan(&star)
	fmt.Printf("请输入终止页（>=start） :")
	fmt.Scan(&end)

	Dowrk(star,end)
}
