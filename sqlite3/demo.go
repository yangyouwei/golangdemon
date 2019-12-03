package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var SqlDB *sql.DB

type Line struct {
	Id      int    `json:"rowid"`
	Ipaddr  string `json:"ipaddr"`
	Comment string `json:"comment"`
}

//create database

//CREATE TABLE "lines"(
//[ipaddr] CHAR(100) NOT NULL,
//[comment] CHAR(100) NOT NULL);
//
//CREATE TABLE [system](
//[contry] CHAR(100) NOT NULL,
//[speedmod] CHAR(100) NOT NULL,
//[getcontryurl] CHAR(200) NOT NULL,
//[getlineurl] CHAR(200) NOT NULL,
//[speedcontry] CHAR(200) NOT NULL);


//初始化方法
func init() {
	var err error
	SqlDB, err = sql.Open("sqlite3", "./sqlite3.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	//连接检测
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	//show all lines
	a,b :=selectline()
	if b != nil {
		fmt.Println(b)
		return
	}
	fmt.Println(a)

	//to json
	//ljson, err := json.Marshal(a)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(ljson))

	l := Line{}

	//add line
	//l.Ipaddr = "192.168.10.100:443"
	//l.Comment = "新加坡"
	//fmt.Printf("add line %v \n",l.Ipaddr)
	//rs,err := AddLine(l)
	//if err != nil {
	//	fmt.Println(rs)
	//}
	//c,d :=selectline()
	//if d != nil {
	//	fmt.Println(d)
	//}
	//fmt.Println(c)

	//delete line
	//l.Id = 6
	//DeleteLine(l)
	//c,d :=selectline()
	//if d != nil {
	//	fmt.Println(d)
	//}
	//fmt.Println(c)

	//update line
	l.Id = 4
	l.Ipaddr = "255.255.255.255:8000"
	l.Comment = "英国"
	UpdateLine(l)
	c,d :=selectline()
	if d != nil {
		fmt.Println(d)
	}
	fmt.Println(c)

}

func selectline()(lines []Line, err error)  {
	lines = make([]Line, 0)
	rows, err := SqlDB.Query("SELECT rowid, * FROM lines")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for rows.Next() {
		var line Line
		rows.Scan(&line.Id, &line.Ipaddr, &line.Comment)
		lines = append(lines, line)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func UpdateLine(l Line) (ra int64, err error) {
	stmt, err := SqlDB.Prepare("UPDATE lines SET ipaddr=?, comment=? WHERE rowid=?")
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	rs, err := stmt.Exec(l.Ipaddr, l.Comment,l.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	ra, err = rs.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func DeleteLine(l Line) (ra int64, err error) {
	rs, err := SqlDB.Exec("DELETE FROM lines WHERE rowid=?", l.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func AddLine(l Line) (id int64, err error) {
	rs, err := SqlDB.Exec("INSERT INTO lines(ipaddr,comment) VALUES (?, ?)", l.Ipaddr,l.Comment)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err = rs.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	return
}
