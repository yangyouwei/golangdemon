package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Line struct {
	Ipaddr  string `json:"ip_port"`
	Comment string `json:"comment"`
}

type Lines struct {
	Lines []Line `json:"lines"`
}

func main()  {
	var l Lines
	l.Lines = append(l.Lines,Line{Ipaddr:"192.168.2.100：443",Comment:"美国"})
	l.Lines = append(l.Lines,Line{Ipaddr:"192.168.2.100：443",Comment:"美国"})
	b,err := json.Marshal(l)
	if err != nil {
		fmt.Println(err)
	}
	var str bytes.Buffer
	_ = json.Indent(&str, []byte(b), "", "	")
	fmt.Println(str.String())
}


// func main()  {
// 	var l Lines
// 	l.Lines = append(l.Lines,Line{Ipaddr:"192.168.2.100：443",Comment:"美国"})
// 	b,err := json.MarshalIndent(l,"","	")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(b))
// }
