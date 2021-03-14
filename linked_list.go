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

func (l *LinkedList) Add(el interface{}) {
	panic("implement me")
}

func (l *LinkedList) Insert(index int, el interface{}) {
	panic("implement me")
}

func (l *LinkedList) Get(index int) interface{} {
	panic("implement me")
}

func (l *LinkedList) Find(el interface{}) (index int, ok bool) {
	panic("implement me")
}

func (l *LinkedList) Remove(index int) {
	panic("implement me")
}

func (l *LinkedList) Length() int {
	return l.length
}
