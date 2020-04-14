package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type UserInfo struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

var Users = make(map[string]UserInfo)

func main() {
	var serverport string = "8000"
	log.Println("server is listen on 0.0.0.0:8000")
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	//get user info
	s.HandleFunc("/GETUsrInfo", GETUsrInfo).Methods("GET")
	//set user info
	s.HandleFunc("/SETUserInfo_urlargs", SETUserInfo_urlargs).Methods("GET")
	s.HandleFunc("/SETUserInfo_bodykv", SETUserInfo_bodykv).Methods("POST")
	s.HandleFunc("/SETUserInfo_bodyjson", SETUserInfo_bodyjson).Methods("POST")

	err := http.ListenAndServe(":"+serverport, r)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GETUsrInfo(w http.ResponseWriter, r *http.Request) {
	CrossDomain(w, r)
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	name := r.Form.Get("name")
	if name == "" {
		fmt.Println("args err： name is none")
		w.Write([]byte("please give name to serarch user info."))
		return
	}

	if _, ok := Users[name]; ok {
		auser := Users[name]
		auserjson, err := json.Marshal(auser)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(auserjson)
		return
	} else {
		w.Write([]byte("user is not exsit."))
		return
	}
}

func SETUserInfo_urlargs(w http.ResponseWriter, r *http.Request)  {
	CrossDomain(w, r)
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("request url args error"))
		return
	}
	auser := UserInfo{}
	auser.Name = r.Form.Get("name")
	auser.Age,err = strconv.Atoi(r.Form.Get("age"))
	auser.Address = r.Form.Get("address")
	Users[auser.Name] = auser
	w.Write([]byte("user set is ok"))
}

//post kv
func SETUserInfo_bodykv(w http.ResponseWriter, r *http.Request) {
	CrossDomain(w, r)
	auser := UserInfo{}
	auser.Name = r.PostFormValue("name")
	age, err := strconv.Atoi(r.PostFormValue("age"))
	if err != nil {
		fmt.Println(err)
		return
	}
	a,_ := ioutil.ReadAll(r.Body)
	fmt.Println(string(a))

	auser.Age = age
	auser.Address = r.PostFormValue("address")
	Users[auser.Name] = auser
	w.Write([]byte("user set is ok"))
}

//post json
func SETUserInfo_bodyjson(w http.ResponseWriter, r *http.Request) {
	CrossDomain(w, r)
	userjson, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	auser := UserInfo{}
	err = json.Unmarshal(userjson, &auser)
	if err != nil {
		fmt.Println(err)
	}
	Users[auser.Name] = auser
	w.Write([]byte("user set is ok"))
}

func CrossDomain(w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	log.Println("request domain ", r.Host, "URL: ", r.URL)
	w.Header().Set("Access-Control-Allow-Origin", "*")                                                                                              //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Connection, User-Agent, Cookie,Action, Module") //header的类型
	w.Header().Set("content-type", "application/json")                                                                                              //返回数据格式是json
	//header("Access-Control-Allow-Credentials : true");
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	return w
}
