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
		l.addFirst(el)
		return
	}
	l.Insert(l.length, el)
}

func (l *LinkedList) addFirst(el interface{}) {
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
		l.addFirst(el)
		return
	}

	var prev *Node
	if index == l.length {
		prev = l.last
	} else {
		prev = l.getNode(index - 1)
	}

	node := &Node{
		el:   el,
		next: prev.next,
	}
	prev.next = node

	if l.length == index {
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
	if l.length == 0 {
		return 0, false
	}
	var n = l.first
	for i := 0; n.next != nil; n = n.next {
		if el == n.el {
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
	var prev *Node
	if index == l.length-1 {
		prev = l.last
	} else {
		prev = l.getNode(index - 1)
	}

	prev.next = prev.next.next
	if l.length-1 == index {
		l.last = prev
	}
	l.length--
}

func (l *LinkedList) Length() int {
	return l.length
}

func (l *LinkedList) Copy() List {
	list := &LinkedList{}
	for i := 0; i < l.length; i++ {
		list.Add(l.Get(i))
	}
	return list
}
