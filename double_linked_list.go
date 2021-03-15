package ta_lab3

import "fmt"

type BiDirectionalNode struct {
	el   interface{}
	next *BiDirectionalNode
	prev *BiDirectionalNode
}

type DoubleLinkedList struct {
	firstSentinel *BiDirectionalNode
	last          *BiDirectionalNode
	length        int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	sentinelNode := &BiDirectionalNode{}
	return &DoubleLinkedList{sentinelNode, sentinelNode, 0}
}

func (node *BiDirectionalNode) PrintBiNode() {
	fmt.Printf("%v, %T| %v, %v\n", node.el, node.el, node.next, node.prev)
}

func (l *DoubleLinkedList) PrintDoubleLinkedList() {
	node := l.firstSentinel
	for node.next != nil {
		node = node.next
		node.PrintBiNode()
	}
	fmt.Println()
}

func NewBiDirectionalNode(el interface{}) *BiDirectionalNode {
	return &BiDirectionalNode{
		el:   el,
		next: nil,
		prev: nil,
	}
}

func (l *DoubleLinkedList) Add(el interface{}) {
	newNode := NewBiDirectionalNode(el)

	lastNode := l.Last()
	lastNode.next = newNode
	newNode.prev = lastNode

	l.last = newNode
	l.length++
}

func (l *DoubleLinkedList) Insert(index int, el interface{}) {
	if index >= l.Length() {
		l.Add(el)
	} else {
		node := l.GetNode(index)
		newNode := NewBiDirectionalNode(el)

		node.prev.next = newNode
		newNode.prev = node.prev
		node.prev = newNode
		newNode.next = node

		l.length++
	}

}

func (l *DoubleLinkedList) Get(index int) interface{} {
	return l.GetNode(index).el
}

func (l *DoubleLinkedList) GetNode(index int) *BiDirectionalNode {
	if index >= l.Length() || index < 0 {
		panic("index out of bounds")
	}

	node := l.First()
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (l *DoubleLinkedList) Find(el interface{}) (index int, ok bool) {
	node := l.firstSentinel

	for i := 0; i < l.Length(); i++ {
		node = node.next
		if node.el == el {
			return i, true
		}
	}
	return -1, false
}

func (l *DoubleLinkedList) Remove(index int) {
	node := l.GetNode(index)
	if index == l.Length()-1 {
		l.last = node.prev
	} else {
		node.next.prev = node.prev
	}

	node.prev.next = node.next
	l.length--
}

func (l *DoubleLinkedList) Length() int {
	return l.length
}

func (l *DoubleLinkedList) First() *BiDirectionalNode {
	if l.firstSentinel.next == nil {
		return l.firstSentinel
	} else {
		return l.firstSentinel.next
	}

}

func (l *DoubleLinkedList) Last() *BiDirectionalNode {
	return l.last
}
