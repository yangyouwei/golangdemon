package main

import (
	"log"
	"os"
	"text/template"
)

type Service struct {
	Name string
	Address string
	Port string
}

type Services struct {
	Name string
	Services []Service
}


const templateText =`{{.}}
upstream {{.Name}} {
	zone upstream-{{.Name}} 64k;{{range .Services}}
	server {{.Address}}:{{.Port}} max_fails=3 fail_timeout=60 weight=1;{{else}}server 127.0.0.1:65535; # force a 502{{end}}
	ip_hash;
}`

func main()  {
	s1 := Service{"aaa","192.168.1.1","8000"}
	s2 := Service{"bbb","192.168.1.100","9000"}
	s3 := Service{"ccc","192.168.1.200","7000"}

	s := Services{}
	s.Name = "anzhi_api"
	s.Services = append(s.Services,s1,s2,s3)

	t := template.Must(template.New("upstream").Parse(templateText))
	err := t.Execute(os.Stdout, s)
	if err != nil {
		log.Println("Executing template:", err)
	}
}
