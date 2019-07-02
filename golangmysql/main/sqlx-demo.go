package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

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

func init() {
	database, err := sqlx.Open("mysql", "root:Yang@123@tcp(192.168.1.28:3306)/web")
	if err != nil {
		fmt.Println("open mysql failed ,", err)
		return
	}
	Db = database
}

//alter
//func main()  {
//	r,err := Db.Exec("alter ")
//}

//inster
//func main()  {
//	r,err := Db.Exec("insert into person(username,sex,email)values(?,?,?)","student","man","student@qq.com")
//	if err != nil {
//		fmt.Println("exec failed ",err)
//		return
//	}
//	id, err := r.LastInsertId()
//	if err != nil {
//		fmt.Println("exec failed ",err)
//		return
//	}
//
//	fmt.Println("insert success ", id)
//}

//select
func main() {
	c := []Chapter{}
	err := Db.Select(&c, "select content from chapter where id=?", 3)
	if err != nil {
		fmt.Println("exec failed ", err)
	}
	fmt.Println(c[0].Content)
}

//delete
//func main()  {
//	_,err := Db.Exec("delete from chapter where id=?",2)
//	if err !=nil {
//		fmt.Println("exec failed",err)
//		return
//	}
//	fmt.Println("delete success")
//}
