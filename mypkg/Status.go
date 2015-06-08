package mypkg

type Status struct {
	id    int
	State string
	Next  map[int][]*Status
}

func (s *Status) AddNext(on int, tar *Status) (ret bool, info string) {
	s.Next[on] = append(s.Next[on], tar)
	return true, ""
}
