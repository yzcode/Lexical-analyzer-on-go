package main

import (
	"bufio"
	"fmt"
	"github.com/yzcode/LAOG/mypkg"
	"io"
	"os"
	"strings"
)

func main() {
	reg_file := "lex.txt"
	src_file := "test.in"
	for i, val := range os.Args {
		if i == 1 {
			reg_file = val
		} else if i == 2 {
			src_file = val
		}
	}
	reg_exp_str := make([]string, 0)
	if fin, err := os.Open(reg_file); err == nil {
		defer fin.Close()
		fin := bufio.NewReader(fin)
		for {
			if instr, inerr := fin.ReadString('\n'); inerr != io.EOF {
				instr = strings.Replace(instr, "\r\n", "", 1)
				//fmt.Println(instr)
				reg_exp_str = append(reg_exp_str, instr)
			} else {
				break
			}
		}
	} else {
		fmt.Println(err)
	}
	run_reg := mypkg.RegGram{}
	run_reg.Build(reg_exp_str)
	run_nfa := mypkg.NFA{}
	run_src := make([]string, 0)
	if isok, info := run_nfa.Build(&run_reg); isok == false {
		fmt.Println(info)
	} else {
		//fmt.Println(run_nfa.States[1].Next['i'][2].State)
		if sfin, err := os.Open(src_file); err == nil {
			defer sfin.Close()
			sfin := bufio.NewReader(sfin)
			for {
				if instr, inerr := sfin.ReadString('\n'); inerr != io.EOF {
					instr = strings.Replace(instr, "\r\n", "", 1)
					//fmt.Println(instr)
					run_src = append(run_src, instr)
				} else {
					break
				}
			}
		} else {
			fmt.Println(err)
		}
	}
	lex_table_file := "lex_table.txt"
	fout, err := os.Create(lex_table_file)
	defer fout.Close()
	if err != nil {
		fmt.Println(lex_table_file, "err")
	} else {
		run_nfa.RunAna(run_src, &run_reg)
		for _, val := range run_nfa.Tokens {
			fout.WriteString(fmt.Sprintf("%s\t%s\t%s\r\n", val.Type.Des, val.Type.Father, val.Content))
		}
		fout.WriteString("#\tfile_end\t#\n")
	}
	fmt.Println("--lex scanner finish--")
	//	for _,val:=range run_reg.ProdExps{
	//		fmt.Println(val);
	//	}
}
