package ta_lab3

type Node struct {
	el   interface{}
	next *Node
}

type LinkedList struct {
	first  *Node
	last   *Node
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Add(el interface{}) {
	if l.length == 0 {
		l.AddFirst(el)
		return
	}
	l.Insert(l.length, el)
}

func (l *LinkedList) AddFirst(el interface{}) {
	l.first = &Node{
		el:   el,
		next: l.first,
	}
	l.length++
	if l.length == 1 {
		l.last = l.first
	}
}

func (l *LinkedList) Insert(index int, el interface{}) {
	if index == 0 {
		l.AddFirst(el)
		return
	}

	prev := l.getNode(index - 1)
	node := &Node{
		el:   el,
		next: prev.next,
	}
	prev.next = node

	if l.length-1 == index {
		l.last = node
	}
	l.length++
}

func (l *LinkedList) Get(index int) interface{} {
	return l.getNode(index).el
}

func (l *LinkedList) getNode(index int) *Node {
	node := l.first
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (l *LinkedList) Find(el interface{}) (index int, ok bool) {
	for i := 0; i < l.length; i++ {
		if el == l.Get(i) {
			return i, true
		}
	}
	return 0, false
}

func (l *LinkedList) Remove(index int) {
	if l.length <= 0 || index > l.length-1 {
		panic("index out of bounds")
	}
	if index == 0 {
		l.first = l.first.next
		l.length--
		return
	}
	prev := l.getNode(index - 1)
	prev.next = prev.next.next
	if l.length-1 == index {
		l.last = prev
	}
	l.length--
}

func (l *LinkedList) Length() int {
	return l.length
}
