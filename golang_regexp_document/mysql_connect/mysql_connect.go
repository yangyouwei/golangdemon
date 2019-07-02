package mysql_connect

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/yangyouwei/golang_regexp_document/conf_get"
)

var Db *sqlx.DB

//创建章节结构
type Chapter struct {
	Id          int64  `db:"id"`
	BooksId     int    `db:"booksId"`
	ChapterId   int    `db:"chapterId"`
	Size        int    `db:"size"`
	ChapterName string `db:"chapterName"`
	Content     string `db:"content"`
	//Content     sql.NullString `db:"content"`
	Url        string `db:"url"`
	ChapterId5 int    `db:"chapterId5"`
	Done       int    `db:"done"`
	Path       string `db:"path"`
}

//select语句
var selectid_str string = "SELECT * FROM `chapter` WHERE `id`="
//update 语句

//数据库链接字符串
var datasourcename string = conf_get.Mysql_conf_str.Username + ":" + conf_get.Mysql_conf_str.Password + "@tcp(" + conf_get.Mysql_conf_str.Ipaddress + ":" + conf_get.Mysql_conf_str.Port + ")/" + conf_get.Mysql_conf_str.DatabaseName

//初始化数据库链接
func init(){
	database, err := sqlx.Open("mysql", datasourcename)
	if err != nil {
		panic(err)
	}
	Db = database
}

//查找函数
func Select_id_get_row(id int64) string  {
	c := []Chapter{}
	err := Db.Select(&c, "select content from chapter where id=?", id)
	if err != nil {
		fmt.Println("exec failed ", err)
	}

	n := len(c)
	if n == 0 {
		fmt.Printf("id : %s is none",)
		return ""
	}
	return c[0].Content
}

//获取存储路径
func  Select_id_get_path(id int64) string  {
	c := []Chapter{}
	err := Db.Select(&c, "select path from chapter where id=?",id)
	if err != nil {
		fmt.Println("exec failed ", err)
	}
	n := len(c)
	if n == 0 {
		fmt.Printf("id : %s is none",)
		return ""
	}
	return c[0].Path
}

