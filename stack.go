package ta_lab3

type Stack struct {
	list List
}

func NewStack(list List) *Stack {
	return &Stack{list: list}
}

func (s *Stack) Pop() interface{} {
	lastEl := s.list.Get(s.list.Length() - 1)
	s.list.Remove(s.list.Length() - 1)
	return lastEl
}

func (s *Stack) Push(l interface{}) {
	s.list.Add(l)
}
