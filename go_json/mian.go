package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	//字段名必须是大写否和jeson不能解析。不能被外部访问到。
	Name string `json:"student_name"` //json编码的时候，替换字段名称
	Age  int
}

// go 语言自带了json库，但是效率低  encoding/json 标准库
//github.com/pquerna/ffjson/ffjson  效率高用法和go自带的库是一样的。
func main() {
	//对数组类型编码
	x := [5]int{1, 2, 3, 4, 5}
	s, err := json.Marshal(x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(s))

	//对map类型编码
	m := make(map[string]float64)
	m["zhangsan"] = 100
	mj, err2 := json.Marshal(m)
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println(string(mj))

	//对对象进行编码
	student := Student{"zhangsan", 26}
	sj, err3 := json.Marshal(student)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(string(sj))

	//----------解码
	//对sj进行解码
	var unjson interface{}
	//json.Unmarshal(sj,&unjson)
	json.Unmarshal([]byte(sj), &unjson)
	fmt.Printf("%v", unjson)

}
