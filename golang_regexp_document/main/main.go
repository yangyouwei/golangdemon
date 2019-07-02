package main

import (
	"fmt"
	"github.com/yangyouwei/golang_regexp_document/conf_get"
	"github.com/yangyouwei/golang_regexp_document/mysql_connect"
	"github.com/yangyouwei/golang_regexp_document/regexp_rule"
)


func main(){
	//获取其实id
	id_range  := conf_get.Id_conf_str
	//查找数据库
	select_id_row := mysql_connect.Select_id_get_row(id_range.Sstart_id)
	//正则匹配
	//标签替换为 缩进或换行
	a := regexp_rule.Regexp_repelace_all(select_id_row)
	//删除带标签的行
	b := regexp_rule.Regexp_del_all(a)
	fmt.Println(b)
	//获取存储路径
	s := mysql_connect.Select_id_get_path(id_range.Sstart_id)
	fmt.Println(s)
}