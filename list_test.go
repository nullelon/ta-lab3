package ta_lab3

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func fillWithRandom(list List) {
	for i := 0; i < 1000; i++ {
		list.Add(rand.Int())
	}
}

func TestLists(t *testing.T) {
	benchList(NewLinkedList(), "LinkedList")
	fmt.Println()

	benchList(NewDoubleLinkedList(), "DoubleLinkedList")
}

func benchList(list List, listName string) {
	fillWithRandom(list)
	middleEl := list.Get(list.Length() / 2)

	iterations := 1000

	var total int64
	for i := 0; i < iterations; i++ {
		l := list.Copy()
		t := time.Now()
		l.Find(middleEl)
		total += time.Since(t).Nanoseconds()

	}
	fmt.Println("find \t" + listName + "\t" + strconv.FormatInt(total/int64(iterations), 10) + "ms")

	total = 0
	for i := 0; i < iterations; i++ {
		l := list.Copy()
		t := time.Now()
		l.Insert(0, 123)
		total += time.Since(t).Nanoseconds()

	}
	fmt.Println("insertFront \t" + listName + "\t" + strconv.FormatInt(total/int64(iterations), 10) + "ms")

	total = 0
	for i := 0; i < iterations; i++ {
		l := list.Copy()
		t := time.Now()
		l.Insert(l.Length()/2, 123)
		total += time.Since(t).Nanoseconds()

	}
	fmt.Println("insertAgn \t" + listName + "\t" + strconv.FormatInt(total/int64(iterations), 10) + "ms")

	total = 0
	for i := 0; i < iterations; i++ {
		l := list.Copy()
		t := time.Now()
		l.Insert(l.Length(), 123)
		total += time.Since(t).Nanoseconds()

	}
	fmt.Println("insertBack \t" + listName + "\t" + strconv.FormatInt(total/int64(iterations), 10) + "ms")
	total = 0
	for i := 0; i < iterations; i++ {
		l := list.Copy()
		t := time.Now()
		l.Remove(0)
		total += time.Since(t).Nanoseconds()

	}
	fmt.Println("deletionFront \t" + listName + "\t" + strconv.FormatInt(total/int64(iterations), 10) + "ms")

	total = 0
	for i := 0; i < iterations; i++ {
		l := list.Copy()
		t := time.Now()
		l.Remove(l.Length() / 2)
		total += time.Since(t).Nanoseconds()

	}
	fmt.Println("deletionAgn \t" + listName + "\t" + strconv.FormatInt(total/int64(iterations), 10) + "ms")

	total = 0
	for i := 0; i < iterations; i++ {
		l := list.Copy()
		t := time.Now()
		l.Remove(l.Length() - 1)
		total += time.Since(t).Nanoseconds()

	}
	fmt.Println("deletionBack\t" + listName + "\t" + strconv.FormatInt(total/int64(iterations), 10) + "ms")

}

func testList(t *testing.T, l List) {
	l.Add(123)
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)

	if l.Get(0) != 123 {
		t.Fatalf("Expected 123, got %v", l.Get(0))
	}

	if l.Get(1) != 1 {
		t.Fatalf("Expected 1, got %v", l.Get(1))
	}

	if l.Get(2) != 2 {
		t.Fatalf("Expected 2, got %v", l.Get(2))
	}

	l.Remove(1)

	if l.Get(0) != 123 {
		t.Fatalf("Expected 123, got %v", l.Get(0))
	}
	if l.Get(1) != 2 {
		t.Fatalf("Expected 2, got %v", l.Get(1))
	}

	l.Remove(0)

	if l.Get(0) != 2 {
		t.Fatalf("Expected 2, got %v", l.Get(0))
	}

	l.Insert(2, 555)
	if l.Get(2) != 555 {
		t.Fatalf("Expected 555, got %v", l.Get(2))
	}

	if n, ok := l.Find(555); !ok || n != 2 {
		t.Fatalf("Expected 2, got %v", n)
	}

}

func TestLinkedList(t *testing.T) {
	testList(t, NewLinkedList())
}

func TestDoubleLinkedList(t *testing.T) {
	testList(t, NewDoubleLinkedList())
}
