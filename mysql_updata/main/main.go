package main

import (
	"github.com/yangyouwei/golang_regexp_document/conf_get"
	"github.com/yangyouwei/golang_regexp_document/mysql_connect"
)


func main(){
	//获取其实id
	IdRange  := conf_get.Id_conf_str
	idchan := make(chan int64,20)
	go func() {
		for i := IdRange.Sstart_id; i <= IdRange.End_id ; i++  {
			idchan <- i
		}
	}()

	for {
		i := <- idchan
		if i == IdRange.End_id {
			break
		}
		mysql_connect.Select_id_get_lenth(i)
	}
}

//select * from chapter order by id desc limit 1
