package mypkg

import (
	//"fmt"
	"fmt"
	"regexp"
	"strings"
)

type NFA struct {
	Start      int
	StatusMap  map[int]Status
	_StatusMap map[string]int
	State      *Set
	States     []Status
	Tokens     []Token
	alpha      map[int]int
}

func (nfa *NFA) Addnode(s string) *Status {
	nfa.States = append(nfa.States, Status{len(nfa.States), s, map[int][]*Status{}})
	if len(s) > 0 {
		nfa._StatusMap[s] = len(nfa.States) - 1
	}
	return &nfa.States[len(nfa.States)-1]
}
func (nfa *NFA) Build(s *RegGram) (ret bool, info string) {
	nfa.alpha = make(map[int]int)
	for i := ' '; i <= '~'; i++ {
		nfa.alpha[int(i)] = 1
	}
	nfa._StatusMap = make(map[string]int)
	nfa.StatusMap = make(map[int]Status)
	nfa.States = make([]Status, 1)
	info = ""
	ret = true
	nfa.Start = nfa.Addnode("__start__").id
	//fmt.Printf("%v %v\n",s.End["id"],s.End["123"])
	//fmt.Printf("%v\n",nfa.States[1].State)
	for _, val := range s.ProdExps {
		left_exp := val.left
		right_exp := strings.Split(val.right, " ")
		regs := regexp.MustCompile(`"(.*)"`)
		nt_exp := string(regs.FindSubmatch([]byte(val.right))[1])
		var left_node, right_node *Status
		if nfa._StatusMap[left_exp] == 0 {
			left_node = nfa.Addnode(left_exp)
		} else {
			left_node = &nfa.States[nfa._StatusMap[left_exp]]
		}
		right_node = &nfa.States[nfa.Start]
		if len(right_exp) >= 2 {
			if nfa._StatusMap[right_exp[0]] == 0 {
				right_node = nfa.Addnode(right_exp[0])
			} else {
				right_node = &nfa.States[nfa._StatusMap[right_exp[0]]]
			}
		} else if len(right_exp) == 0 {
			return false, fmt.Sprintf("regular expression: %v (right : %v ) is illegal the len of exp is %v\n", val, right_exp, len(right_exp))
		}
		tmp_src := right_node
		for i, val := range nt_exp {
			if i == len(nt_exp)-1 {
				break
			}
			tmp_node := nfa.Addnode("")
			tmp_src.AddNext(int(val), tmp_node)
			tmp_src = tmp_node
		}
		tmp_src.AddNext(int(nt_exp[len(nt_exp)-1]), left_node)
	}
	fmt.Printf("NFA has been built successful\n")
	return
}
func (nfa *NFA) Token_nfa(src string, pos int, rg *RegGram) (int, *Set, string) {
	//fmt.Printf("s: %s ||| pos: %d\n",src,pos)
	//	if src[pos]=='"'{
	//		fmt.Printf("fuck off\n")
	//		fmt.Printf("%s\n",src)
	//	}
	nfa.State = NewSet()
	nfa.State.Add(nfa.Start)
	endpos := pos
	for i := pos; i < len(src); i++ {
		//		if src[pos]=='"'{
		//			fmt.Printf("i: %d ---------\n",i)
		//		}
		//
		chr := src[i]
		tmp_states := NewSet()
		//fmt.Printf("--- %c ---",src[i])
		for _, val := range nfa.State.List() {
			if len(nfa.States[val].Next[int(chr)]) != 0 {

				for _, value := range nfa.States[val].Next[int(chr)] {
					tmp_states.Add(value.id)
				}
			}
		}
		//		if src[pos]=='"'{
		//			for _,val:= range tmp_states.List(){
		//				fmt.Printf("%v ",nfa.States[val].Next['1'][0].State);
		//			}
		//			fmt.Println()
		//		}
		//fmt.Printf("\n")
		if tmp_states.IsEmpty() {
			break
		}
		endpos = i
		nfa.State = tmp_states
	}
	final_set := NewSet()
	for _, val := range nfa.State.List() {
		if rg.End[nfa.States[val].State] != 0 {
			final_set.Add(nfa.States[val].id)
		}
	}
	return endpos + 1, final_set, string([]rune(src)[pos : endpos+1])
}
func (nfa *NFA) RunAna(src []string, rg *RegGram) (ret bool, info string) {
	//fmt.Println(len(src))
	lexical_error := false
	for i, val := range src {
		//fmt.Printf("line number is %d\n",i+1)
		lineNumber := i + 1
		pos := 0
		for pos < len(val) && !lexical_error {
			for pos < len(val) && val[pos] == ' ' || val[pos] == '\t' {
				pos++
			}
			if pos < len(val) {
				//fmt.Printf("pos: %d\n",pos)
				posx, finalset, token := nfa.Token_nfa(val, pos, rg)

				pos = posx
				//fmt.Printf("pos: %d\n",pos)
				if finalset.IsEmpty() {
					lexical_error = true
					fmt.Printf("lexical error at line %d, column %d\n", lineNumber, pos)
					break
				}
				for _, val := range TokenTypeSets.TokenTypes {
					if finalset.Has(nfa._StatusMap[val.Des]) {
						nfa.Tokens = append(nfa.Tokens, Token{TokenTypeSets.GetType(val.Des), token, lineNumber, pos})
						break
					}
				}
			}
		}
	}
	return true, ""
}
