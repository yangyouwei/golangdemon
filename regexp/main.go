package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)


//爬取流浪地球  影片名称和评分
func main()  {
	url := "https://movie.douban.com/subject/26266893/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doubanhtml, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(doubanhtml))
	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*)</span>`)   //分组，第一个分组是全部匹配的结果，第二个是括号里的。
	result := reg.FindAllStringSubmatch(string(doubanhtml),-1)  //使用for循环然后取切片的下标，或者使用result1[0][1]直接取出
	for _, v := range result {
		fmt.Println(v[1])
	}

	//fmt.Println(string(doubanhtml))
	regs := regexp.MustCompile(`<strong\s*class="ll rating_num" property="v:average">(.*)</strong>`)
	result1 := regs.FindAllStringSubmatch(string(doubanhtml),-1)
	for _, v := range result1 {
		fmt.Println(v[1])
	}
}

//常用函数
//func ()  {
//	regexp.Match()
//	regexp.MatchString()
//	下面两个一起使用
//	a := regexp.MustCompile()
//	a.FindAllString()
//}


//显示匹配上的字符
//func main() {
//	reg := regexp.MustCompile("yyw")
//	result := reg.FindAllString("yywlfsadfasdf",-1)
//	fmt.Printf("%v \n",result)
//
//	reg1 := regexp.MustCompile(`^z(.*)1$`)
//	result1 := reg1.FindAllString("zkdkifef432451",-1)
//	fmt.Printf("%v \n",result1)
//	}

//判断是否匹配上 true or false
//func main()  {
//	isok , err := regexp.Match("[a-zA-Z]{3}",[]byte("y23"))
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(isok)
//}
//执行返回 false
