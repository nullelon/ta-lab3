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
