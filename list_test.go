package ta_lab3

import (
	"math/rand"
	"testing"
)

func fillWithRandom(list List) {
	for i := 0; i < 10000; i++ {
		list.Add(rand.Int())
	}
}

func doBench(b *testing.B, list List) {
	fillWithRandom(list)
	middleEl := list.Get(list.Length() / 2)
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	b.Run("indexing", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			list.Find(middleEl)
		}
	})
	b.Run("insertionFront", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			list.Insert(0, 0)
		}
	})
	b.Run("insertionAgn", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			list.Insert(list.Length()/2, 0)
		}
	})
	b.Run("insertionEnd", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			list.Insert(list.Length()-1, 0)
		}
	})
	b.Run("deletionFront", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			list.Remove(0)
		}
	})
	b.Run("deletionAgn", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			list.Find(list.Length() / 2)
		}
	})
	b.Run("deletionEnd", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			list.Find(list.Length() - 1)
		}
	})
}

func BenchmarkDoubleLinkedList(b *testing.B) {
	doBench(b, NewDoubleLinkedList())
}

func BenchmarkLinkedList(b *testing.B) {
	doBench(b, NewLinkedList())
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
