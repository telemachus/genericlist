package genericlist_test

import (
	"container/list"
	"fmt"
	"io"
	"testing"

	"github.com/telemachus/genericlist"
)

func BenchmarkStdlibPushFront(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := list.New()
		for j := 0; j < 1000; j++ {
			l.PushFront(j)
		}
	}
}

func BenchmarkGenericPushFront(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := genericlist.New[int]()
		for j := 0; j < 1000; j++ {
			l.PushFront(j)
		}
	}
}

func BenchmarkStdlibPushBack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := list.New()
		for j := 0; j < 1000; j++ {
			l.PushBack(j)
		}
	}
}

func BenchmarkGenericPushBack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := genericlist.New[int]()
		for j := 0; j < 1000; j++ {
			l.PushBack(j)
		}
	}
}

func BenchmarkStdlibRemove(b *testing.B) {
	l := list.New()
	for j := 0; j < 1000; j++ {
		l.PushBack(j)
	}
	es := make([]*list.Element, 0, 1000)
	for e := l.Front(); e != nil; e = e.Next() {
		es = append(es, e)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range es {
			v := l.Remove(e)
			fmt.Fprintf(io.Discard, "%d\n", v.(int))
		}
	}
}

func BenchmarkGenericRemove(b *testing.B) {
	l := genericlist.New[int]()
	for j := 0; j < 1000; j++ {
		l.PushBack(j)
	}
	es := make([]*genericlist.Element[int], 0, 1000)
	for e := l.Front(); e != nil; e = e.Next() {
		es = append(es, e)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range es {
			v := l.Remove(e)
			fmt.Fprintf(io.Discard, "%d\n", v)
		}
	}
}

func BenchmarkStdlibInsertBefore(b *testing.B) {
	l := list.New()
	for j := 0; j < 1000; j++ {
		l.PushBack(j)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e := l.Front().Next()
		for j := 0; j < 100; j++ {
			l.InsertBefore(j, e)
		}
	}
}

func BenchmarkGenericInsertBefore(b *testing.B) {
	l := genericlist.New[int]()
	for j := 0; j < 1000; j++ {
		l.PushBack(j)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e := l.Front().Next()
		for j := 0; j < 100; j++ {
			l.InsertBefore(j, e)
		}
	}
}

func BenchmarkStdlibInsertAfter(b *testing.B) {
	l := list.New()
	for j := 0; j < 1000; j++ {
		l.PushBack(j)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e := l.Front().Next()
		for j := 0; j < 100; j++ {
			l.InsertAfter(j, e)
		}
	}
}

func BenchmarkGenericInsertAfter(b *testing.B) {
	l := genericlist.New[int]()
	for j := 0; j < 1000; j++ {
		l.PushBack(j)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e := l.Front().Next()
		for j := 0; j < 100; j++ {
			l.InsertAfter(j, e)
		}
	}
}

func BenchmarkStdlibMoveToFront(b *testing.B) {
	l := list.New()
	for j := 0; j < 1000; j++ {
		l.PushBack(j)
	}
	es := make([]*list.Element, 0, 1000)
	for e := l.Front(); e != nil; e = e.Next() {
		es = append(es, e)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range es {
			l.MoveToFront(e)
		}
	}
}

func BenchmarkGenericMoveToFront(b *testing.B) {
	l := genericlist.New[int]()
	for j := 0; j < 1000; j++ {
		l.PushBack(j)
	}
	es := make([]*genericlist.Element[int], 0, 1000)
	for e := l.Front(); e != nil; e = e.Next() {
		es = append(es, e)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range es {
			l.MoveToFront(e)
		}
	}
}

func BenchmarkStdlibMoveToBack(b *testing.B) {
	l := list.New()
	for j := 0; j < 1000; j++ {
		l.PushFront(j)
	}
	es := make([]*list.Element, 0, 1000)
	for e := l.Front(); e != nil; e = e.Next() {
		es = append(es, e)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range es {
			l.MoveToBack(e)
		}
	}
}

func BenchmarkGenericMoveToBack(b *testing.B) {
	l := genericlist.New[int]()
	for j := 0; j < 1000; j++ {
		l.PushFront(j)
	}
	es := make([]*genericlist.Element[int], 0, 1000)
	for e := l.Front(); e != nil; e = e.Next() {
		es = append(es, e)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range es {
			l.MoveToBack(e)
		}
	}
}

func BenchmarkStdlibMoveBefore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := list.New()
		e := l.PushFront(1)
		for i := 0; i < 100; i++ {
			l.MoveBefore(e, &list.Element{Value: i})
		}
	}
}

func BenchmarkGenericMoveBefore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := genericlist.New[int]()
		e := l.PushFront(1)
		for i := 0; i < 100; i++ {
			l.MoveBefore(e, &genericlist.Element[int]{Value: i})
		}
	}
}

func BenchmarkStdlibMoveAfter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := list.New()
		e := l.PushFront(1)
		for i := 0; i < 100; i++ {
			l.MoveAfter(e, &list.Element{Value: i})
		}
	}
}

func BenchmarkGenericMoveAfter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := genericlist.New[int]()
		e := l.PushFront(1)
		for i := 0; i < 100; i++ {
			l.MoveAfter(e, &genericlist.Element[int]{Value: i})
		}
	}
}

func BenchmarkStdlibPushBackList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := list.New()
		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3)
		l.PushBackList(l)
	}
}

func BenchmarkGenericPushBackList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := genericlist.New[int]()
		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3)
		l.PushBackList(l)
	}
}

func BenchmarkStdlibPushFrontList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := list.New()
		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3)
		l.PushFrontList(l)
	}
}

func BenchmarkGenericPushFrontList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := genericlist.New[int]()
		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3)
		l.PushFrontList(l)
	}
}
