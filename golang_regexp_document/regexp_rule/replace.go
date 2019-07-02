package regexp_rule

import "regexp"

//正则替换函数
func Regexp_repelace_all(s string) string {
	//replace <br />
	reg := regexp.MustCompile("<br />")
	result := reg.ReplaceAllString(s,"\n")

	//replace &nbsp;&nbsp;&nbsp;&nbsp;
	reg1 := regexp.MustCompile("&nbsp;&nbsp;&nbsp;&nbsp;")
	result1 := reg1.ReplaceAllString(result,"    ")

	return result1
}

func Regexp_del_all(s string) string {
	//replace 删除带标签的行
	reg := regexp.MustCompile(".*\\<[\\S\\s]+?\\>.*")
	result := reg.ReplaceAllString(s,"")
	return result
}

