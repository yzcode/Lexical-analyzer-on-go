package mypkg

import (
	"fmt"
	"strings"
)

type RegGram struct {
	ProdExps     []ProdExp
	ProdExpsMap  map[ProdExp]int
	_ProdExpsMap map[int]ProdExp
	Vn           map[string]int
	Vt           map[string]int
	End          map[string]int
}

func (reggram *RegGram) Build(inputs []string) (ret bool, info string) {
	reggram.End = make(map[string]int)
	reggram.Vn = make(map[string]int)
	reggram.Vt = make(map[string]int)
	reggram._ProdExpsMap = make(map[int]ProdExp)
	reggram.ProdExpsMap = make(map[ProdExp]int)
	ret = true
	for i, input := range inputs {
		//fmt.Println(i);
		if input[0] == '#' || len(input) == 0 {
			continue
		}
		if i == 0 {
			TokenTypeSets.Build(input, reggram.End)
			continue
		}
		var tmp_input = strings.Split(input, " = ")
		if len(tmp_input) > 2 {
			return false, "one RegGram shouldn't have two ="
		}
		reggram.ProdExps = append(reggram.ProdExps, ProdExp{tmp_input[0], tmp_input[1]})
	}
	for _,val:=range reggram.ProdExps{
		fmt.Println(val)
	}
	fmt.Printf("%v lines Regular expression has been read\n", len(reggram.ProdExps))
	return
}
