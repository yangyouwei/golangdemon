package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	query()
	//query2()
	//insert()
	//update()
	//remove()
}



//查询数据
func query() {
	db, err := sql.Open("mysql", "root:@/shopvisit")
	check(err)

	rows, err := db.Query("SELECT * FROM shopvisit.announcement")
	check(err)

	for rows.Next() {
		columns, _ := rows.Columns()

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		//将数据保存到 record 字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
	rows.Close()

}
func query2()  {
	fmt.Println("Query2")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/shopvisit?charset=utf8")
	check(err)

	rows, err := db.Query("SELECT id,imgUrl,createDate,state FROM announcement")
	check(err)

	for rows.Next(){
		var id int
		var state int
		var imgUrl string
		var createDate string
		//注意这里的Scan括号中的参数顺序，和 SELECT 的字段顺序要保持一致。
		if err := rows.Scan(&id,&imgUrl,&createDate,&state); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s id is %d on %s with state %d\n", imgUrl, id, createDate, state)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
}


//插入数据
func insert()  {
	db, err := sql.Open("mysql", "root:@/shopvisit")
	check(err)

	stmt, err := db.Prepare(`INSERT announcement (imgUrl, detailUrl, createDate, state) VALUES (?, ?, ?, ?)`)
	check(err)

	res, err := stmt.Exec("/visitshop/img/ann/cofox1.png",nil,"2017-09-06",0)
	check(err)

	id, err := res.LastInsertId()  //必须是自增id的才可以正确返回。
	check(err)

	fmt.Println(id)
	stmt.Close()

}

//修改数据
func update() {
	db, err := sql.Open("mysql", "root:@/shopvisit")
	check(err)

	stmt, err := db.Prepare("UPDATE announcement set imgUrl=?, detailUrl=?, createDate=?, state=? WHERE id=?")
	check(err)

	res, err := stmt.Exec("/visitshop/img/ann/cofox2.png", nil, "2017-09-05", 1, 7)
	check(err)

	num, err := res.RowsAffected()  //返回值是1 修改成功，如果是零则是不成功（修改值和原来的值一样，mysql实际没有修改，或者没有这条记录）
	check(err)

	fmt.Println(num)
	stmt.Close()
}

//删除数据
func remove() {
	db, err := sql.Open("mysql", "root:@/shopvisit")
	check(err)

	stmt, err := db.Prepare("DELETE FROM announcement WHERE id=?")//问号是占位符
	check(err)

	res, err := stmt.Exec(7)  //7插入到上面的问号中
	check(err)

	num, err := res.RowsAffected()
	check(err)

	fmt.Println(num)
	stmt.Close()

}

func check(err error) {
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
}


//mysql 需要安装驱动 https://github.com/Go-SQL-Driver/MySQL，只导入init方法即可。_ "github.com/go-sql-drive/mysql"
//必须同时加载database/sql 这个包
//sql.Open() 打开一个注册过的数据库驱动
