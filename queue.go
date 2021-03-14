package ta_lab3

type Queue struct {
	list List
}

func NewQueue(list List) *Queue {
	return &Queue{list: list}
}

func (s *Queue) Pop() interface{} {
	el := s.list.Get(0)
	s.list.Remove(0)
	return el
}

func (s *Queue) Push(l interface{}) {
	s.list.Add(l)
}
