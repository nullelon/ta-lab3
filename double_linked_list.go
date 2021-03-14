package ta_lab3

type DoubleLinkedList struct {
	first  *Node
	length int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (l *DoubleLinkedList) Add(el interface{}) {
	panic("implement me")
}

func (l *DoubleLinkedList) Insert(index int, el interface{}) {
	panic("implement me")
}

func (l *DoubleLinkedList) Get(index int) interface{} {
	panic("implement me")
}

func (l *DoubleLinkedList) Find(el interface{}) (index int, ok bool) {
	panic("implement me")
}

func (l *DoubleLinkedList) Remove(index int) {
	panic("implement me")
}

func (l *DoubleLinkedList) Length() int {
	return l.length
}
