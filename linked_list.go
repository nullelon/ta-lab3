package ta_lab3

type Node struct {
	el   interface{}
	next *Node
}

type LinkedList struct {
	first  *Node
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) add(el interface{}) {
	panic("implement me")
}

func (l *LinkedList) put(el interface{}, index int) {
	panic("implement me")
}

func (l *LinkedList) get(index int) interface{} {
	panic("implement me")
}

func (l *LinkedList) find(el interface{}) (index int, ok bool) {
	panic("implement me")
}

func (l *LinkedList) remove(index int) {
	panic("implement me")
}
