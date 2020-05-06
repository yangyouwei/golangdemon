package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
	"strings"
	"time"
)

var (
	//mysql
	//MysqlString string = "root" + ":" + "Yang@1008" + "@tcp(" + "192.168.1.202" + ":" + "3306" + ")/" + "shop"
	MysqlString string = "root" + ":" + "GaoPeng@123" + "@tcp(" + "192.168.2.5" + ":" + "3306" + ")/" + "yyw-shop"
	//mysql 连接池
	Db *sql.DB
	//err
	err error
)

type ManagerStr struct {
	Id       int    `json:"mg_id"`
	Rid      int    `json:"role_id"`
	UserName string `json:"mg_name"`
	Password string `json:"mg_pwd"`
	Regdate  int    `json:"mg_time"`
	Mobile   string `json:"mg_mobile"`
	Email    string `json:"mg_email"`
	State    int    `json:"mg_state"`
}

func init()  {
	Db, err = sql.Open("mysql",MysqlString)
	check(err)

	err = Db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}


func check(err error) {
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}
}


func CountRows(db *sql.DB)  {
	//statement := "SELECT COUNT(*) FROM `sp_manager`;"
	statement := fmt.Sprintf("SELECT * FROM `sp_manager` WHERE `mg_name` = \"%s\";","admin")
	rows, err := db.Query(statement)
	if	err != nil{
		log.Println(err)
	}
	var user ManagerStr
	for rows.Next(){
		//注意这里的Scan括号中的参数顺序，和 SELECT 的字段顺序要保持一致。

		if err := rows.Scan(&user.Id,&user.UserName,&user.Password,&user.Regdate,&user.Rid,&user.Mobile,&user.Email,&user.State); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(user)
}

func getusers(db *sql.DB,startnum int,endnum int)  {
	//　select * from table limit n-1,m-n;
	statement := fmt.Sprintf("select * `from sp_manager` limit %v,%v;",startnum-1,endnum-startnum)
	rows, err := db.Query(statement)
	if	err != nil{
		log.Println(err)
	}

	var users []ManagerStr

	for rows.Next(){
		//注意这里的Scan括号中的参数顺序，和 SELECT 的字段顺序要保持一致。
		var user ManagerStr
		if err := rows.Scan(&user.Id,&user.UserName,&user.Password,&user.Regdate,&user.Rid,&user.Mobile,&user.Email,&user.State); err != nil {
			log.Fatal(err)
		}
		users = append(users,user)
	}
	fmt.Println(users)
}

func adduser(db *sql.DB)  {
	stmt, err := db.Prepare(`INSERT INTO sp_manager (mg_name,mg_pwd,mg_time,mg_mobile,mg_email) VALUES (?,?,?,?,?)`)
	check(err)
	defer stmt.Close()
	res, err := stmt.Exec("admin100","aslkdjf2343t45gedrg",time.Now().Unix(),13333333,"admin@163.com")
	check(err)
	id, err := res.LastInsertId()  //必须是自增id的才可以正确返回。
	check(err)
	log.Println(id)
}

func main()  {
	adduser(Db)
	a := ManagerStr{}

	t := reflect.TypeOf(a)
	fieldNum := t.NumField()
	result := make([]string,0,fieldNum)
	for i:= 0;i<fieldNum;i++ {
		tagName := t.Field(i).Name
		tags := strings.Split(string(t.Field(i).Tag), "\"")
		if len(tags) > 1 {
			tagName = tags[1]
		}
		result = append(result,tagName)
	}
	fmt.Println(result)
}
