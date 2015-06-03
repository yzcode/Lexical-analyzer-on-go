package mypkg

import "fmt"

type ProdExp struct {
	left  string
	right string
}

func (p ProdExp) String() string {
	return fmt.Sprintf("%v -> %v ", p.left, p.right)
}
