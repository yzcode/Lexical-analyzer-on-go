package main

import (
	"regexp"
	"fmt"
)

func main() {
	str :="\" \"\""
	regs := regexp.MustCompile(`"(.*)"`)
	nt_exp := regs.FindSubmatch([]byte(str))
	for _,val:=range nt_exp{
		fmt.Println(string(val))
	}
}
