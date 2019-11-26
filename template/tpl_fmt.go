package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type Mgs struct {
	ScopeId      int    `json:"scopeId"`
	Name         string `json:"name"`
	Id0          int    `json:"id0"`
	Id1          int    `json:"id1"`
	AlarmMessage string `json:"alarmMessage"`
	StartTime    int    `json:"startTime"`
}

const tmpl  = `ScopeId = {{.ScopeId}}
Name = {{.Name}}
Id0 = {{.Id0}}
Id1 = {{.Id1}}
StartTime = {{.StartTime}}
AlarmMessage = {{.AlarmMessage}}`


func main() {
	sweaters := Mgs{ScopeId:2,Name:"test",Id1:0,Id0:3,AlarmMessage:"testsetsetsetset",StartTime:234245623}

	tmpl, err := template.New("alarm").Parse(tmpl)  //建立一个模板
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, sweaters) 
	if err != nil {
		panic(err)
	}
	fmt.Println(buf)
}
