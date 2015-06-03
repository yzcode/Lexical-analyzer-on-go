package mypkg

import (
	"fmt"
	"strings"
)

type TokenType struct {
	Des    string
	Id     int
	Father string
}
type TokenTypeSet struct {
	TokenTypes []TokenType
	size       int
}

var TokenTypeSets TokenTypeSet

func (p *TokenTypeSet) Build(input string, ends map[string]int) (ret bool, info string) {
	p.size = 0
	for _, val := range strings.Split(input, " ") {
		tmp_str := strings.Split(val, ":")
		if len(tmp_str) > 2 {
			return false, "one TokenType should not have two father"
		} else if len(tmp_str) == 1 {
			p.TokenTypes = append(p.TokenTypes, TokenType{tmp_str[0], p.size, ""})
			fmt.Println(tmp_str[0])
		} else {
			p.TokenTypes = append(p.TokenTypes, TokenType{tmp_str[0], p.size, tmp_str[1]})
		}
		ends[tmp_str[0]] = 1
		p.size++
	}
	//	for _, val := range p.TokenTypes {
	//		fmt.Printf("%v 's id is %v ans father is %s\n",val.Des, val.Id, val.Father)
	//	}
	fmt.Printf("%v kind(s) of Token Types has been read\n", p.size)
	return true, ""
}
func (p *TokenTypeSet) GetType(input string) *TokenType {
	for _, val := range p.TokenTypes {
		if strings.EqualFold(val.Des, input) {
			return &val
		}
	}
	return nil
}

type Token struct {
	Type    *TokenType
	Content string
	Row     int
	Col     int
}
