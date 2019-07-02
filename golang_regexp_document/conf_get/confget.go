package conf_get

import (
	"github.com/Unknwon/goconfig"
	"log"
)

//mysql配置参数结构
type Mysql_conf struct {
	Username string
	Password string
	Ipaddress string
	Port	string
	DatabaseName	string
}

//id 结构
type Id struct {
	Sstart_id int64
	End_id int64
}

//正则表达式  替换
type Regexp_list_replace struct {
	rep_replace_map map[string]string
}

//正则表达式  删除
type Regexp_del struct {
	rep_del_map map[string]string
}

//创建解析出来参数变量
var Mysql_conf_str Mysql_conf
var Id_conf_str Id
var Regexp_replace_str Regexp_list_replace
var Regexp_del_str Regexp_del
var Store_path string

//初始化配置文件读取，初始化变量
func init()  {
	cfg, err := goconfig.LoadConfigFile("conf")
	if err != nil {
		log.Println("读取配置文件失败[config.ini]")
		panic(err)
	}

		Mysql_fun(cfg,err)
		id_fun(cfg,err)
		regexp_del_fun(cfg,err)
		regexp_replace_fun(cfg,err)
}

//mysql配置参数 初始化
func Mysql_fun(c *goconfig.ConfigFile,err error) {
	m := &Mysql_conf_str
	m.Username, err = c.GetValue("mysql", "username")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "username", err)
		panic(err)
	}

	m.Password, err = c.GetValue("mysql", "password")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "password", err)
		panic(err)
	}

	m.Ipaddress, err = c.GetValue("mysql", "ipaddr")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "ipaddr", err)
		panic(err)
	}

	m.Port, err = c.GetValue("mysql", "prot")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "prot", err)
		panic(err)
	}

	m.DatabaseName, err = c.GetValue("mysql", "database")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "database", err)
		panic(err)
	}
}
//id配置参数 初始化
func id_fun(c *goconfig.ConfigFile,err error)  {
	id := &Id_conf_str
	id.Sstart_id, err = c.Int64("id", "start")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "start", err)
		panic(err)
	}

	id.End_id, err = c.Int64("id", "end")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "end", err)
		panic(err)
	}
}

//正则表达式 参数初始化
func regexp_del_fun(c *goconfig.ConfigFile,err error)  {

}

//正则表达式 参数初始化
func regexp_replace_fun(c *goconfig.ConfigFile,err error)  {

}

//存储参数初始化
func store_path(c *goconfig.ConfigFile,err error)  {
	sp := &Store_path
	*sp ,err = c.GetValue("store", "path")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "path", err)
		panic(err)
	}
}