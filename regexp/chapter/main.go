package main

import (
	"bufio"
	"fmt"
	"github.com/Unknwon/goconfig"
	"io"
	"log"
	"os"
	"regexp"
)

type Chapter1 struct {
	rules []string
}
type Chapter2 struct {
	rules []string
}

var Chapterrules1 Chapter1
var Chapterrules2 Chapter2

func init() {
	cfg, err := goconfig.LoadConfigFile("conf")
	if err != nil {
		log.Println("读取配置文件失败[config.ini]")
		panic(err)
	}
	Chapterrules1.getchapterrules(cfg)
	Chapterrules2.getchapterrules2(cfg)
}

func main() {
	var fp string = "test.txt"
	readfile(fp)
}

func readfile(fp string) {
	fi, err := os.Open(fp)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		matchrules(string(a))
	}
}

func matchrules(s string) {
	//headertxt := `\s*.*(序言|楔子|文案|引子|序|序章|自序|前言|起始|前传|序幕)\s*：*`
	rules := `\s*.*(序言|楔子|文案|引子|序|序章|自序|前言|起始|前传|序幕)\s*：*`
	isok, err := regexp.Match(rules, []byte(s))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(s,"---",rules,":",isok)
	if isok {
		fmt.Println(s, "\t\t", rules)
	}
}

func (this *Chapter1) getchapterrules(c *goconfig.ConfigFile) {
	this.rules = c.GetKeyList("chapter_rules1")

}

func (this *Chapter2) getchapterrules2(c *goconfig.ConfigFile) {
	this.rules = c.GetKeyList("chapter_rules2")
}
