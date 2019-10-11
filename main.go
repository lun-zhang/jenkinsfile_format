package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	filename string
	replace  bool
	indent   string
)

func init() {
	flag.StringVar(&filename, "fn", "", "Jenkinsfile文件")
	flag.BoolVar(&replace, "r", false, "是否替换原来的Jenkinsfile文件，默认不替换而是输出到屏幕")
	flag.StringVar(&indent, "idt", "	", "补齐用的字符串，默认是一个制表符")
	flag.Parse()
}

func main() {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	formatted := format(string(bs), indent)
	if !replace {
		fmt.Println(formatted)
	} else {
		if err := ioutil.WriteFile(filename, []byte(formatted), os.ModePerm); err != nil {
			panic(err)
		}
	}
}

func format(original string, indent string) (formatted string) {
	if original == "" || indent == "" {
		return original
	}
	lines := strings.Split(original, "\n")

	indents := ""
	hasBlockCommon := false

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		//FIXME: 需支持任意块注释
		//以"/*"开头且以"*/"结尾的块注释不格式化
		if n := len(line); n >= 2 {
			if !hasBlockCommon && line[:2] == "/*" {
				hasBlockCommon = true
				formatted += line + "\n"
				continue
			}
			if hasBlockCommon && line[n-2:] == "*/" {
				hasBlockCommon = false
				formatted += line + "\n"
				continue
			}
		}
		if hasBlockCommon {
			formatted += line + "\n"
			continue
		}

		noComment := line
		if ci := strings.Index(line, "//"); ci > 0 {
			noComment = strings.Trim(noComment[:ci], " ")
		}
		if noComment == "" {
			continue
		}
		end := noComment[len(noComment)-1]
		if end == '}' {
			indents = indents[:len(indents)-len(indent)]
			formatted += indents + line + "\n"
			continue
		}
		formatted += indents + line + "\n"
		if end == '{' {
			indents = indents + indent
		}
	}
	return
}
